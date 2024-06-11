Hey what's going on everyone, this is Bek Brace channel, my name is Amir, good morning good evening wherever you maybe.

Writing a currency converter application and dockerizing it with chainguard Go image:

# STEP1
# Create a new directory for your project
mkdir currency_converter
cd currency_converter

# Initialize a new Go module
go mod init currency_converter

# Install the Cobra library, which is used for creating the CLI
go get github.com/spf13/cobra@latest


# STEP2
# Use your preferred text editor to create the main.go file
vscode main.go
then the program itself with comments

# Step 3: Build the Go Application
go build -o currency_converter

# Step 4: Test the Go Application
# Run your application to ensure it works
./currency_converter convert 100 USD EUR

# Step 5: Dockerizing the Application with Chainguard
Create a dokerfile

