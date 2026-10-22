# MichaelMylettComp415Project2

The GO project file is divided into 5 functions, each one completing one step of the project. Note that for all times a "median" is provided, if the list is an even length, that the median is only the latter of the two values at the middle of the list, rather than the average of those two values. All lists used for finding a median are sorted using slices.Sort(). Also note that for color printing to work, the project must be run using "go run main.go" in the terminal, rather than using the run button. The program uses custom structs to sort out every aspect of a specific job, often using lists, or "slices" as a method for storing many jobs. This document uses the word slice/list interchangeably, but they are referring to the same thing. The program requires no user input, so long as no files are modified.

# First

The function finds every job that is considered entry level and from the year 2024 and adds the salaries to a total sum, dividing it by the number of job listings meeting these requirements to find an average. The same is then done for the year 2023. If the 2024 average is less than the 2023 average, it prints both average salaries in red. Otherwise, it prints the same message in green.

# Second

The function takes all entry level jobs from 2024, puts their salaries in a list, and finds the median of that list. It then does the same for all entry level jobs from 2023. Then, it prints the difference in median salaries for the two years. If the 2023 median is greater, this prints in red, otherwise it prints in green.
The function then finds the average salary for all jobs from the year 2024, as well as the median, using the same methods described above. It then prints whether the median or average was greater, and the difference between them.

# Third

The function finds all jobs from 2024, putting their salaries in one list if they are entry level or another if mid level (if neither of these, they are not added to any list). It then compares the median salary of the entry level and mid level jobs, printing the difference between the two as an expected raise if you were to move up from an entry level job to a mid level one. This is printed in green. If the mid level median is somehow equal to or lower than the entry level median, it will state so in red, along with the difference between the medians.

# Fourth

The function takes all 2024 jobs, putting their salaries into 4 different lists based on the following criteria:
      - fully remote, entry level
      - fully onsite, entry level
      - fully remote, all levels
      - fully onsite, all levels
It then records the median of each list and prints them to the terminal, stating whether it is for a remote/onsite or entry-level/any-level job. For both entry level and all level jobs, it prints whether the remote or on site median salary was better, with it printing in red if the on site median salary is higher or blue if otherwise.

# Fifth

The function runs through every job listing from 2024 that does not have "research" in the title (not case sensitive), forming a map where the job titles are used as keys and slices of the salaries are used as values. The function appends new salary values to the slices as it increments through the file. A for loop takes every item in the map previously mentioned and adds the medians of each job to a new map as values, with the keys being the same. It then uses this second map to find the top five median salaries, doing so using a series of 5 for loops that compare the median values with each other. It then outputs the top median values along with what job they were for in the console. The highest value appears in blue, with the rest being green.

# Another important note: the csv file being read must be called "salaries.csv", as the project is hard-coded to only accept that file name. This may be case sensitive on some systems, and not case sensitive on others.
