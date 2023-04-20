# TFS to PDF

Application that can pull details about work item in TFS and save it as PDF. 
Output name will be the same as TFS item ID. 
You must provide TFS item ID in parameter to this application


## Technical details

This application using [chromedp](https://github.com/chromedp/chromedp) as an engine to convert HTML page to PDF. The template for HTML located in `\template` folder and filled using GO [html/template](https://pkg.go.dev/html/template) rules


## Deployment

Once main application complied the deployment layout should be:

```
template\
  |- template_files\
  \- template.htm

tfs2pdf.exe
```

* Some path to folder such as `\bin` are hardcoded in app itself