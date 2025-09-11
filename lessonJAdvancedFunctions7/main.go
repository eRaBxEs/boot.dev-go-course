package main

/*
Assignment
Run the code without changing anything: you should see a compilation error.
Fix the scoping issue in the function so that it runs as you'd expect.
*/

func splitEmail(email string) (string, string) {

	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}
