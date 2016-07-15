

{{ define "content" }}

<tr><td class="title">Invitation</td></tr>
<tr><td>Invitation code: <b>{{ .InviteCode}}</b></td></tr>
<tr><td>Please <a href="{{ .InviteLink}}">click here</a> to signup at TargetSeq.</td></tr>

{{ end }}
