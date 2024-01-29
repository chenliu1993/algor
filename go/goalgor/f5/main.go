package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// type Vpm struct {
// 	PrivateNIC string `yaml:"PrivateNIC"`
// 	InsideNIC  string `yaml:"InsideNIC,omitempty"`
// }

type AWS struct {
	ImageID []string `yaml:"imageId"`
}

type HardwareType struct {
	//Vpm Vpm `yaml:"Vpm"`
	Aws AWS `yaml:"aws"`
}

type certifiedHardware struct {
	HardwareType map[string]HardwareType `yaml:"certifiedHardware"`
}

// type certifiedHardware struct {
// 	HardwareType map[string]struct {
// 		Aws *struct {
// 			ImageID []string `yaml:"imageId"`
// 		} `yaml:"aws"`
// 	} `yaml:"certifiedHardware"`
// }

// download the yaml and parse them
func getHarewareType(url string) (*certifiedHardware, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	hardwareTypes := &certifiedHardware{}
	err = yaml.Unmarshal(content, hardwareTypes)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return hardwareTypes, nil
}

func getOccurences(certified *certifiedHardware) map[string]int {
	occurences := map[string]int{}
	for _, val := range certified.HardwareType {
		for _, imageId := range val.Aws.ImageID {
			if _, ok := occurences[imageId]; !ok {
				occurences[imageId] = 1
			} else {
				occurences[imageId] = occurences[imageId] + 1
			}
		}
	}
	return occurences
}

func main() {
	// step 1
	hardwareTypes, err := getHarewareType("https://vesio.blob.core.windows.net/releases/certified-hardware/aws.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("The types are")
	fmt.Println(hardwareTypes)

	occurences := getOccurences(hardwareTypes)
	fmt.Println(occurences)

}

/* certifiedHardware.HardwareType = map[string]HardwareType{
	"aws1": HardwareType{
		Aws: AWS{
			ImageID: []string{
				"1",
				"2",
			},
		},
	},
	"aws2": HardwareType{
		Aws: AWS{
			ImageID: []string{
				"31",
				"23",
			},
		},
	},
}

file, err := os.Create("aws1.yaml")
if err != nil {
	panic(err)
}

out, err := yaml.Marshal(certifiedHardware)
if err != nil {
	panic(err)
}
_, err = file.Write(out)
if err != nil {
	panic(err)
} */

/* content, err := os.ReadFile("./aws.yaml")
if err != nil {
	panic(err)
}
err = yaml.Unmarshal(content, certifiedHardware)
if err != nil {
	panic(err)
}
fmt.Println(certifiedHardware) */
