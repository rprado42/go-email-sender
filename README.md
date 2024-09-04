# Go Email Sender

This project is a Go application that reads email addresses from an Excel file and sends a welcome email to each address.

## Project Structure

- `main.go`: Contains the core functionality to read emails from Excel and send welcome emails.
- `data/emails.xlsx`: Example Excel file containing email addresses (update with your own).
- `README.md`: Project overview and setup instructions.
- `go.mod` and `go.sum`: Go module files for dependency management.

## Setup

1. **Install Dependencies**

   Make sure you have Go installed, then run:

   ```bash
   go mod tidy


2. **Update Configuration**

Open main.go and update the SMTP server configuration:
smtpHost: Your SMTP server address
smtpPort: Your SMTP server port (e.g., 587)
smtpUser: Your email address
smtpPass: Your email password

3. **Add Emails**

Place your emails.xlsx file in the data/ directory. Ensure that the file has email addresses in the first column of the first sheet.

4. **Run the Application**

go run main.go



### **Additional Notes**

1. **`LICENSE` File**: Ensure you create a `LICENSE` file in your project directory with the full text of the MIT License or whichever license you choose. The `README.md` file will reference this `LICENSE` file for details.

2. **Placing the License Section**: The License section is generally placed at the end of the `README.md` file, after the setup instructions and other relevant sections, to keep the document organized and user-friendly. 

By following this structure, your `README.md` will provide clear, organized instructions for setting up and using your project, along with licensing information at the end.



### 3. **`go.mod` and `go.sum`**

Generate these files by running the following command in your project directory:

```bash
go mod init go-email-sender


go get github.com/xuri/excelize/v2
go get gopkg.in/gomail.v2
