package demo

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const form2 = `
<html>
	<body>
		<h1>Statistics</h1>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body>
</html>
`

const table = `
	<html>
	<body>
		<h1>Statistics</h1>
		<table>
			<thead>
				<tr>
					<th>Results</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Numbers</td>
					<td>%v</td>
				</tr>
				<tr>
					<td>Count</td>
					<td>%f</td>
				</tr>
				<tr>
					<td>Mean</td>
					<td>%f</td>
				</tr>
				<tr>
					<td>Median</td>
					<td>%f</td>
				</tr>
			</tbody>
		</table>
	</body>
	</html>
`

func InitStatistics() {
	http.HandleFunc("/", statisticsServer)
	err := http.ListenAndServe("localhost:8000", nil)
	checkError(err)
}

func statisticsServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, form2)
	case "POST":
		{
			text := req.FormValue("in")
			arr := resolveStringToFloatArr(text)

			fmt.Printf("value %#v", arr)
			// io.WriteString(w, "123")
			fmt.Fprintf(w, table, arr, sum1(arr), mean1(arr), median(arr))
		}
	}
}

func median(arr []float64) float64 {
	sort.Float64s(arr)
	lens := len(arr)
	if lens%2 == 0 {
		mindIndex1 := lens / 2
		mindIndex2 := mindIndex1 - 1
		return (arr[mindIndex1] + arr[mindIndex2]) / 2
	} else {
		mindIndex := lens / 2
		return arr[mindIndex]
	}
}

func mean1(arr []float64) (res float64) {
	sum := sum1(arr)
	return sum / float64(len(arr))
}

func sum1(arr []float64) (res float64) {
	for _, v := range arr {
		res += v
	}
	return
}

func resolveStringToFloatArr(text string) (result []float64) {
	arr := strings.Split(text, " ")
	for _, v := range arr {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			result = append(result, f)
		}
	}
	return
}
