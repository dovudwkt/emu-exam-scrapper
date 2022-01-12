# emu-exam-scrapper

scrapper for emu exam time table

## Endpoints

- http://localhost:3030/exams/import - imports the data from the external source and saves as a JSON file.
- http://localhost:3030/exams/search?courses=xxx,yyy - returns the exam dates for provided courses. courses should be comma seperated course codes without spaces in between.

Example: `http://localhost:3030/exams/search?courses=ITEC122,ITEC114,ITEC397`

<hr>

Application for Ubuntu: [emu-exam-scrapper](https://github.com/dovudwkt/emu-exam-scrapper/raw/main/emu-exam-scrapper) 

Application for Windows: [emu-exam-scrapper.exe](https://github.com/dovudwkt/emu-exam-scrapper/raw/main/emu-exam-scrapper.exe)
