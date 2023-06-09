# TFS to PDF

Application that pull details about work item in TFS and save it as PDF. 
Output name will be the same as TFS item ID. 
You must provide TFS item ID in parameter to this application


## Technical details

This application using [chromedp](https://github.com/chromedp/chromedp) as an engine to convert HTML page to PDF. The template for HTML located in `\templates` folder and filled using GO [html/template](https://pkg.go.dev/html/template) rules - so it must exist. The path to `\templates` folder is hardcoded into application

## How to run

If you building this application on your machine:

1. Just clone repository: `git clone git@github.com:AF250329/tfs2pdf.git`
2. Execute `go build .` inside cloned folder
3. Open any shell (for example Powershell) and run executable with parameters:

```pwsh
tfs2pdf.exe --token=<your TFS token> <TFS item ID>
```

for example:

```pwsh
tfs2pdf.exe --token=yx4pcq7a3erlwrbxxxosio4az7aoatxqa2qs65fub  318371
```

The PDF file will be generated with the same name as work item ID.

### Notes
- Read here: [Wiki: How to generate TFS token](https://github.com/AF250329/tfs2pdf/wiki/How-to-create-TFS-access-token) in case you need to generate TFS token
- In case images are missing in PDF and you see only image place holders - Please disconnect from Pulse Secure and try again

## Deployment

Once main application complied the deployment layout should be:

```
templates\
  |- template_files\
  \- template.htm

tfs2pdf.exe
```
<br /><br />
![go logo](./.github/go.logo.small.png)
