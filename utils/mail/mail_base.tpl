
{{define "base"}}
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" dir="ltr"> 
<head>
<style type="text/css">
*{
    font-family:'Segoe UI', 'Helvetica Neue Medium', Arial, sans-serif; 
}
td.prefix {
    font-size:14px; 
    color:#707070;
}
td.title{
    padding-bottom: 25px;
    font-size:18px;
    color:#2672ec;
}
td.signature{
    padding-top: 25px;
    font-size:14px;
    color:#2a2a2a;
}
td{
    font-size:14px;
    color:#2a2a2a;
}
a{
    color:#2672ec; 
    text-decoration:underline;
}
</style>

<title></title>
</head>
<body>
<table dir="ltr">
      <tr><td class="prefix">iGeneTech Account</td></tr>
      {{template "content" .}}
      <tr><td class="signature"> </td></tr>
      <tr><td>Best</td></tr>
      <tr><td>iGeneTech Team</td></tr>
</table>
</body>
</html>
{{end}}
