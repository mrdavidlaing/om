package download_clients

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"log"
	"regexp"
	"sort"
	"strings"
)

//counterfeiter:generate -o ./fakes/product_versioner.go --fake-name ProductVersioner . productVersioner
type productVersioner interface {
	GetAllProductVersions(slug string) ([]string, error)
}

func DetermineProductVersion(
	slug string,
	exactVersion string,
	versionRegex string,
	versioner productVersioner,
	stderr *log.Logger,
) (string, error) {
	productVersions, err := versioner.GetAllProductVersions(slug)
	if err != nil {
		return "", err
	}

	existingVersions := strings.Join(productVersions, ", ")
	if existingVersions == "" {
		existingVersions = "none"
	}

	if versionRegex != "" {
		foundVersion, err := findLatestVersionFromRegex(productVersions, versionRegex, stderr)
		if err != nil {
			return "", fmt.Errorf("no valid versions found for product %q and product version regex %q\nexisting versions: %s\n%w", slug, versionRegex, existingVersions, err)
		}
		return foundVersion, nil
	}

	if exactVersion != "" {
		for _, version := range productVersions {
			if version == exactVersion {
				return exactVersion, nil
			}
		}
	}
	return "", fmt.Errorf("no valid versions found for product %q and product version %q\nexisting versions: %s", slug, exactVersion, existingVersions)
}

func findLatestVersionFromRegex(productVersions []string, regex string, stderr *log.Logger) (string ,error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("could not compile regex %q: %w", regex, err)
	}

	var versions version.Collection
	for _, productVersion := range productVersions {
		if !re.MatchString(productVersion) {
			continue
		}

		v, err := version.NewVersion(productVersion)
		if err != nil {
			stderr.Printf(fmt.Sprintf("warning: could not parse semver version from: %s", productVersion))
			continue
		}
		versions = append(versions, v)
	}

	sort.Sort(versions)

	if len(versions) == 0 {
		return "", fmt.Errorf("no available version found")
	}

	return versions[len(versions)-1].Original(), nil
}