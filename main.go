package main

const host = "https://localhost:8043/api/v2"
const user = "omada-tf-test"
const password = "V8qMh%k%J*xq&87aco"

//func main() {
//
//	client := &http.Client{}
//
//	requestURL := fmt.Sprintf("%s/%s", host, "login")
//	fmt.Printf("ATTEMPTING TO LOG IN\n")
//	jsonBody := []byte(fmt.Sprintf(`{"username": %s, "password": %s}`, user, password))
//	bodyReader := bytes.NewReader(jsonBody)
//
//	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
//
//	if err != nil {
//		fmt.Printf(fmt.Sprintf("ERROR 1::: %s", err.Error()))
//		os.Exit(1)
//	}
//
//	req.Header.Add("Content-Type", "application/json")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Printf(fmt.Sprintf("ERROR 2::: %s", err.Error()))
//		os.Exit(1)
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Printf(fmt.Sprintf("ERROR 3::: %s", err.Error()))
//		}
//	}(resp.Body)
//	body, err := io.ReadAll(resp.Body)
//
//	fmt.Printf(string(body))
//}

func main() {

}
