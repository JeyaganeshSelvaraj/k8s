package types

type ConfigMap struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	DataToAdd map[string]string `json:"dataToAdd"`
	DataToDel []string          `json:"dataToDel"`
	DataToUpd map[string]string `json:"dataToUpd"`
}
