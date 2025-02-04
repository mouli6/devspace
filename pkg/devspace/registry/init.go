package registry

import (
	"github.com/devspace-cloud/devspace/pkg/devspace/docker"
	"github.com/pkg/errors"

	"github.com/devspace-cloud/devspace/pkg/devspace/config/versions/latest"
	"github.com/devspace-cloud/devspace/pkg/devspace/kubectl"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreatePullSecrets creates the image pull secrets
func CreatePullSecrets(config *latest.Config, client *kubectl.Client, dockerClient docker.ClientInterface, log log.Logger) error {
	if config.Images != nil {
		pullSecrets := []string{}
		createPullSecrets := map[string]bool{}

		for _, imageConf := range config.Images {
			if imageConf.CreatePullSecret == nil || *imageConf.CreatePullSecret == true {
				registryURL, err := GetRegistryFromImageName(imageConf.Image)
				if err != nil {
					return err
				}

				createPullSecrets[registryURL] = true
			}
		}

		for registryURL := range createPullSecrets {
			displayRegistryURL := registryURL
			if displayRegistryURL == "" {
				displayRegistryURL = "hub.docker.com"
			}

			log.StartWait("Creating image pull secret for registry: " + displayRegistryURL)
			err := createPullSecretForRegistry(config, client, dockerClient, registryURL, log)
			log.StopWait()
			if err != nil {
				return errors.Errorf("Failed to create pull secret for registry: %v", err)
			}

			pullSecrets = append(pullSecrets, GetRegistryAuthSecretName(registryURL))
		}

		if len(pullSecrets) > 0 {
			err := addPullSecretsToServiceAccount(client, pullSecrets, log)
			if err != nil {
				return errors.Wrap(err, "add pull secrets to service account")
			}
		}
	}

	return nil
}

func addPullSecretsToServiceAccount(client *kubectl.Client, pullSecrets []string, log log.Logger) error {
	// Get default service account
	serviceaccount, err := client.Client.CoreV1().ServiceAccounts(client.Namespace).Get("default", metav1.GetOptions{})
	if err != nil {
		log.Errorf("Couldn't find service account 'default' in namespace '%s': %v", client.Namespace, err)
		return nil
	}

	// Check if all pull secrets are there
	changed := false
	for _, newPullSecret := range pullSecrets {
		found := false

		for _, pullSecret := range serviceaccount.ImagePullSecrets {
			if pullSecret.Name == newPullSecret {
				found = true
				break
			}
		}

		if found == false {
			changed = true
			serviceaccount.ImagePullSecrets = append(serviceaccount.ImagePullSecrets, v1.LocalObjectReference{Name: newPullSecret})
		}
	}

	// Should we update the service account?
	if changed {
		_, err := client.Client.CoreV1().ServiceAccounts(client.Namespace).Update(serviceaccount)
		if err != nil {
			return errors.Wrap(err, "update service account")
		}
	}

	return nil
}

func createPullSecretForRegistry(config *latest.Config, client *kubectl.Client, dockerClient docker.ClientInterface, registryURL string, log log.Logger) error {
	username, password := "", ""
	if dockerClient != nil {
		authConfig, _ := dockerClient.GetAuthConfig(registryURL, true)
		if authConfig != nil {
			username = authConfig.Username
			password = authConfig.Password
		}
	}

	if config.Deployments != nil && username != "" && password != "" {
		for _, deployConfig := range config.Deployments {
			email := "noreply@devspace.cloud"

			namespace := client.Namespace
			if deployConfig.Namespace != "" {
				namespace = deployConfig.Namespace
			}

			err := CreatePullSecret(client, namespace, registryURL, username, password, email, log)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
