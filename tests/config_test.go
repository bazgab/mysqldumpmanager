package tests

/*
Since go's os.SetEnv method cannot set environment variables globally and permanently,
the process that we have to go through to avoid using a password in the CLI is to:

	1. Pull the user and password from /etc/mysqldumpmanager/conf.yaml
	2. Set these as environment variables for the current shell we are working on
	3. Use the environment variables to push this to /etc/my.cnf and save the file

This way, after initConfig the user will have a /etc/mysqldumpmanager/conf.yaml file configured
with user and password, and also have a /etc/my.cnf file configured with

`
[mysqldump]
user=$MYSQL_USER
password=$MYSQL_PASSWORD
`


*/
