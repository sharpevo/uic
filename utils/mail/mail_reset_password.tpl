
{{ define "content" }}

<tr><td class="title">Password Reset Confirmation</td></tr>
<tr><td>Please <a href="{{ .ResetLink}}">click here</a> to reset the password for the iGeneTech account: <b>{{ .Email}}</b></td></tr>

{{ end }}
