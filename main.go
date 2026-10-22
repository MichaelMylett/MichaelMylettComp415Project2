package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Job struct {
	//job_title,experience_level,employment_type,work_models,work_year,employee_residence,salary,
	//salary_currency,salary_in_usd,company_location,company_size
	job_title          string
	experience_level   string
	employment_type    string
	work_models        string
	work_year          int
	employee_residence string
	salary             int
	salary_currency    string
	salary_in_usd      int
	company_location   string
	company_size       string
}

func first(jobs []Job) {
	salarySum2024 := 0
	counter2024 := 0
	salarySum2023 := 0
	counter2023 := 0
	for _, job := range jobs {
		if job.experience_level == "Entry-level" && job.work_year == 2024 {
			//fmt.Println(job)
			salarySum2024 += job.salary_in_usd
			counter2024 += 1
		}
	}
	for _, job := range jobs {
		if job.experience_level == "Entry-level" && job.work_year == 2023 {
			//fmt.Println(job)
			salarySum2023 += job.salary_in_usd
			counter2023 += 1
		}
	}
	average2024 := salarySum2024 / counter2024
	average2023 := salarySum2023 / counter2023
	//fmt.Println(average2024)
	//fmt.Println(average2023)
	fmt.Println("Part 1:")
	if average2024 < average2023 {
		color.Red("2024 average is %v USD, whereas 2023 average is %v USD.", average2024, average2023)
	} else {
		color.Green("2024 average is %v USD, whereas 2023 average is %v USD.", average2024, average2023)
	}
	fmt.Println()
}
func second(jobs []Job) {
	var salaryList2024 []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.experience_level == "Entry-level" {
			//fmt.Println(job)
			salaryList2024 = append(salaryList2024, job.salary_in_usd)
		}
	}
	listLength2024 := len(salaryList2024)
	slices.Sort(salaryList2024)
	//fmt.Println(salaryList2024)
	median2024 := salaryList2024[listLength2024/2]
	//fmt.Println(median2024) //ask if median with even length should be average of the two values at the center of the list or if it doesn't matter (note that I did the weird one in readme)

	var salaryList2023 []int
	for _, job := range jobs {
		if job.work_year == 2023 && job.experience_level == "Entry-level" {
			//fmt.Println(job)
			salaryList2023 = append(salaryList2023, job.salary_in_usd)
		}
	}
	listLength2023 := len(salaryList2023)
	slices.Sort(salaryList2023)
	//fmt.Println(salaryList2023)
	median2023 := salaryList2023[listLength2023/2]
	//fmt.Println(median2023)
	fmt.Println("Part 2:")
	if median2023 > median2024 {
		difference := median2023 - median2024
		color.Red("The 2023 median is higher, with a difference of %v between 2023 and 2024.", difference)
	} else {
		difference := median2024 - median2023
		color.Green("The 2024 median is higher than or equal to that of 2023, with a difference of %v between 2024 and 2023.", difference)
	}
	//add report difference between average and median, ask if for 2024 or 2023 or both (pick one and note it in the readme)
	salarySum2024 := 0
	counter2024 := 0
	for _, job := range jobs {
		if job.work_year == 2024 {
			//fmt.Println(job)
			salarySum2024 += job.salary_in_usd
			counter2024 += 1
		}
	}
	average2024 := salarySum2024 / counter2024
	//fmt.Println(average2024)

	var salaryList2024new []int
	for _, job := range jobs {
		if job.work_year == 2024 {
			//fmt.Println(job)
			salaryList2024new = append(salaryList2024new, job.salary_in_usd)
		}
	}
	listLength2024new := len(salaryList2024new)
	slices.Sort(salaryList2024new)
	//fmt.Println(salaryList2024new)
	median2024new := salaryList2024new[listLength2024new/2]
	//fmt.Println(median2024new)
	if median2024new > average2024 {
		difference := median2024new - average2024
		fmt.Printf("The 2024 median salary is higher than the 2024 average, with a difference of %v.\n", difference)
	} else {
		difference := average2024 - median2024new
		fmt.Printf("The 2024 average salary is higher than the 2024 median, with a difference of %v.\n", difference)
	}
	fmt.Println()
}
func third(jobs []Job) {
	var salaryList2024entry []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.experience_level == "Entry-level" {
			//fmt.Println(job)
			salaryList2024entry = append(salaryList2024entry, job.salary_in_usd)
		}
	}
	listLength2024entry := len(salaryList2024entry)
	slices.Sort(salaryList2024entry)
	//fmt.Println(salaryList2024entry)

	var salaryList2024med []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.experience_level == "Mid-level" {
			//fmt.Println(job)
			salaryList2024med = append(salaryList2024med, job.salary_in_usd)
		}
	}
	listLength2024med := len(salaryList2024med)
	slices.Sort(salaryList2024med)
	//fmt.Println(salaryList2024med)

	entryMedian2024 := salaryList2024entry[listLength2024entry/2]
	medMedian2024 := salaryList2024med[listLength2024med/2]
	//fmt.Println(entryMedian2024)
	//fmt.Println(medMedian2024)
	fmt.Println("Part 3:")
	if medMedian2024 > entryMedian2024 {
		difference := medMedian2024 - entryMedian2024
		color.Green("You could expect a %v USD raise in 2024 switching from entry level to mid-level based on the median salary of each.", difference)
	} else {
		difference := entryMedian2024 - medMedian2024
		color.Red("Somehow, the median salary for a mid-level job in 2024 is less than the entry level median salary, with a difference of %v USD.", difference)
	}
	fmt.Println()
}
func fourth(jobs []Job) {
	var salaryList2024entryRemote []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.work_models == "Remote" && job.experience_level == "Entry-level" {
			//fmt.Println(job)
			salaryList2024entryRemote = append(salaryList2024entryRemote, job.salary_in_usd)
		}
	}
	listLength2024entryRemote := len(salaryList2024entryRemote)
	slices.Sort(salaryList2024entryRemote)
	//fmt.Println(salaryList2024entryRemote)

	var salaryList2024entryOnsite []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.work_models == "On-site" && job.experience_level == "Entry-level" {
			//fmt.Println(job)
			salaryList2024entryOnsite = append(salaryList2024entryOnsite, job.salary_in_usd)
		}
	}
	listLength2024entryOnsite := len(salaryList2024entryOnsite)
	slices.Sort(salaryList2024entryOnsite)
	//fmt.Println(salaryList2024entryOnsite)

	var salaryList2024allRemote []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.work_models == "Remote" {
			//fmt.Println(job)
			salaryList2024allRemote = append(salaryList2024allRemote, job.salary_in_usd)
		}
	}
	listLength2024allRemote := len(salaryList2024allRemote)
	slices.Sort(salaryList2024allRemote)
	//fmt.Println(salaryList2024allRemote)

	var salaryList2024allOnsite []int
	for _, job := range jobs {
		if job.work_year == 2024 && job.work_models == "On-site" {
			//fmt.Println(job)
			salaryList2024allOnsite = append(salaryList2024allOnsite, job.salary_in_usd)
		}
	}
	listLength2024allOnsite := len(salaryList2024allOnsite)
	slices.Sort(salaryList2024allOnsite)
	//fmt.Println(salaryList2024allOnsite)

	medianEntryRemote := salaryList2024entryRemote[listLength2024entryRemote/2]
	medianEntryOnsite := salaryList2024entryOnsite[listLength2024entryOnsite/2]
	medianAllRemote := salaryList2024allRemote[listLength2024allRemote/2]
	medianAllOnsite := salaryList2024allOnsite[listLength2024allOnsite/2]

	fmt.Println("Part 4:")

	fmt.Println("Remote entry-level job median salary: ", medianEntryRemote)
	fmt.Println("On-site entry-level job median salary: ", medianEntryOnsite)
	fmt.Println("Remote any-level job median salary: ", medianAllRemote)
	fmt.Println("On-site any-level job median salary: ", medianAllOnsite)

	if medianEntryRemote >= medianEntryOnsite {
		difference := medianEntryRemote - medianEntryOnsite
		color.Blue("The median entry-level remote salary is %v USD greater than the median entry-level onsite salary for 2024.\n", difference)
	} else {
		difference := medianEntryOnsite - medianEntryRemote
		color.Red("The median entry-level onsite salary is %v USD greater than the median entry-level remote salary for 2024.\n", difference)
	}

	if medianAllRemote >= medianAllOnsite {
		difference := medianAllRemote - medianAllOnsite
		color.Blue("The median remote salary is %v USD greater than the median onsite salary for 2024.\n", difference)
	} else {
		difference := medianAllOnsite - medianAllRemote
		color.Red("The median onsite salary is %v USD greater than the median remote salary for 2024.\n", difference)
	}
	fmt.Println()
}
func fifth(jobs []Job) {
	nonResearchJobs := make(map[string][]int) //map pairing every job title with list of salaries
	for _, job := range jobs {                //looked up strings.Contains()
		if job.work_year == 2024 && job.experience_level == "Entry-level" && !strings.Contains(job.job_title, "Research") && !strings.Contains(job.job_title, "research") {
			nonResearchJobs[job.job_title] = append(nonResearchJobs[job.job_title], job.salary_in_usd)
		}
	}
	//fmt.Println(nonResearchJobs["Data Engineer"])
	//fmt.Println(nonResearchJobs["Data Scientist"])
	//fmt.Println(nonResearchJobs["Research Analyst"])
	for key, _ := range nonResearchJobs {
		slices.Sort(nonResearchJobs[key])
		//fmt.Println(key, nonResearchJobs[key])
	}
	//fmt.Println(nonResearchJobs["Data Engineer"])
	//fmt.Println(nonResearchJobs["Data Scientist"])

	nonResearchJobsMedians := make(map[string]int)
	for key, value := range nonResearchJobs {
		nonResearchJobsMedians[key] = value[len(value)/2]
	}
	//fmt.Println(nonResearchJobsMedians)

	first := ""
	firstMed := 0
	second := ""
	secondMed := 0
	third := ""
	thirdMed := 0
	fourth := ""
	fourthMed := 0
	fifth := ""
	fifthMed := 0

	for key, value := range nonResearchJobsMedians { //first place
		//fmt.Print(key, value, " ")
		//fmt.Println(value[len(value)/2])
		if value > firstMed {
			first = key
			firstMed = value
		}
	}
	for key, value := range nonResearchJobsMedians { //second place
		if key != first && value > secondMed {
			second = key
			secondMed = value
		}
	}
	for key, value := range nonResearchJobsMedians { //third place
		if key != second && key != first && value > thirdMed {
			third = key
			thirdMed = value
		}
	}
	for key, value := range nonResearchJobsMedians { //fourth place
		if key != third && key != second && key != first && value > fourthMed {
			fourth = key
			fourthMed = value
		}
	}
	for key, value := range nonResearchJobsMedians { //fifth place
		if key != fourth && key != third && key != second && key != first && value > fifthMed {
			fifth = key
			fifthMed = value
		}
	}
	fmt.Println("Part 5:")
	fmt.Println("Top 5 median highest paying entry-level jobs:")
	color.Blue("%v: %v USD\n", first, firstMed)
	color.Green("%v: %v USD\n", second, secondMed)
	color.Green("%v: %v USD\n", third, thirdMed)
	color.Green("%v: %v USD\n", fourth, fourthMed)
	color.Green("%v: %v USD\n", fifth, fifthMed)
}

func main() {
	file, err := os.Open("salaries.csv") //looked up this command and Close() on google
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()               //supposedly unhandled error, doesn't seem to cause problems though and error handling wasn't in the example I saw
	csvreader := csv.NewReader(file) //looked up csv.NewReader() on google

	var jobs []Job

	if _, err := csvreader.Read(); err != nil { //this if statement specifically is entirely copied from a google search (the program worked without it, I just didn't like the minor errors coming up in the terminal [exit code was still 0])
		if errors.Is(err, io.EOF) {
			fmt.Println("The file is empty.")
			return
		}
		log.Fatal(err)
	}

	for { //looked up the first about 7 lines of this for loop on google, but the rest was purely my own code
		record, err := csvreader.Read()
		//fmt.Println(record)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		job_title := record[0]
		experience_level := record[1]
		employment_type := record[2]
		work_models := record[3]
		work_year, err := strconv.Atoi(record[4])
		if err != nil {
			fmt.Println(err)
		}
		employee_residence := record[5]
		salary, err := strconv.Atoi(record[6])
		if err != nil {
			fmt.Println(err)
		}
		salary_currency := record[7]
		salary_in_usd, err := strconv.Atoi(record[8])
		if err != nil {
			fmt.Println(err)
		}
		company_location := record[9]
		company_size := record[10]

		currentJob := Job{
			job_title:          job_title,
			experience_level:   experience_level,
			employment_type:    employment_type,
			work_models:        work_models,
			work_year:          work_year,
			employee_residence: employee_residence,
			salary:             salary,
			salary_currency:    salary_currency,
			salary_in_usd:      salary_in_usd,
			company_location:   company_location,
			company_size:       company_size,
		}
		jobs = append(jobs, currentJob)
	}

	first(jobs)
	second(jobs)
	third(jobs)
	fourth(jobs)
	fifth(jobs)
}
