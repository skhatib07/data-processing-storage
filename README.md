# How to run the code:
Make sure you have `Go 1.23.0` AND `GNU Make >=3.81` installed
If you don't:
    <ul>
    <li>[Install Golang](https://go.dev/doc/install)</li>
    <li>[Install GNU Make](https://www.gnu.org/software/make/#download)</li>
    </ul>

Then either:
    <ol>
    <li>Run `make` in the root directory with the Makefile</li>
    **OR**
    <li>Run `go test ./database -v` in the root directory (this runs the same command as the Makefile)</li>
    </ol>

This command runs the file `database_test.go` in the `./database directory`. This file runs all of the example commands that were demoed in Figure 2 on the [Assignment Document](https://docs.google.com/document/d/1GVdJngVnufjrnXlWGC_eSru_ZxDdOOQkLT7KO2ZHF2Q/edit?tab=t.0#heading=h.4xe0wzkcsuqy).

# How to modify the assigment to become official
I think to modify this assinment to make it official, a few things are needed. Firstly, there should be more explicit requirements on testing the created software. As it stands, there's undefined expectations on whether a testing environment should be created, but in the enterprise software industry, every piece of code must be tested before it is deployed. I think it would be valuable to have the `Assignment Automated Unit Testing` assignment be a sequel to this assignment, where you are expected to implement tests for the database you coded in this assignment. Lastly, I think it would be beneficial for students to develop a light frontend for this assignment. Similar to my previous suggestion, I think having the `React Application` assignment as a sequel to this assignment would be appropriate, where students are instructed to interact with their database through a basic React-based frontend website.
