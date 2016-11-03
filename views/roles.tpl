<div id="ajaxFullSpinner"></div>
<style>
#ajaxFullSpinner {
	position: fixed;
	left: 0px;
	top: 0px;
	width: 100%;
	height: 100%;
	z-index: 9999;
    background: url('/static/img/logo.png') 50% 50% no-repeat rgba(249,249,249,0.7);
	background-size: 300px;
}

.btn{
	margin-right: 5px;
	opacity: .5;
}
.btn-xs{
	line-height: 1;
}
.btn.active:hover,.btn.active,.btn:active {
	opacity: 1;
}

</style>
<table class="table table-hover table-striped">
    <thead>
        <tr>
            <th>Name</th>
            <th>Apps</th>
            <th>Email</th>
            <th>Roles</th>
            <th>UserOperation</th>
        </tr>
    </thead>
    <tbody>
		{{$appList := .AppList}}
        {{range $index, $user := .UserList}}
        <tr data-id="{{$user.Id.Hex}}">
            <td>
                {{$user.Name}}
			</td>
			<td>
				<div class="btn-group" data-toggle="buttons">
                {{range $index, $app := $appList}}
                <label class="btn btn-xs btn-default {{if HasApp $user $app.Id.Hex}}active{{end}} tck" data-userid="{{$user.Id.Hex}}" data-appid="{{$app.Id.Hex}}">
						<input type="checkbox" autocomplete="off" data-userid="{{$user.Id.Hex}}" data-appid="{{$app.Id.Hex}}" data-target="#toggleApp"> {{$app.Name}}
					</label>
				{{end}}
				</div>
            </td>
            <td>{{$user.Email}}</td>
            <td>
                {{range $k, $v := $user.Roles}}
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-rolename="{{$k}}" data-target="#deleteRole"> {{$k}} </a><br/>
                {{end}} 
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-target="#createRole">
                    <span class="glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
                </a>
            </td>
            <td>Edit</td>
        </tr>
        {{end}}
    </tbody>
</table>





<script>
$(window).load(function() {
	ajaxFullSpinner = $("#ajaxFullSpinner")
	$(ajaxFullSpinner).fadeOut(1000);
	$(document).ajaxStart(function(){
		$(ajaxFullSpinner).show()
	}).ajaxComplete(function(){
		$(ajaxFullSpinner).fadeOut(1000);
	});
})

$(document).ready(function(){
    $(".tck").on("click",function(event){
        $userId = $(event.target).data('userid');
        $appId = $(event.target).data('appid');
        console.log($userId)
        console.log($appId)
        $.ajax({
            url: "{{urlfor "UserAppController.Post"}}",
            data: {
                "userId":$userId,
                "appId":$appId,
            },
            method: "POST",
            success: function(response){
                window.location.reload()
            },
            error: function(xhr, status, error){
                alert(xhr.responseText)
            }
        })

    })
})
</script>

<!-- Modal -->
<div class="modal fade" id="createRole" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Add Role</h4>
            </div>
            <form action="{{urlfor "RoleController.Post"}}" method="POST">
                <input type="hidden" name="userId" id="userId" class="form-control" value="" placeholder="Enter the role name">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="roleName" id="roleName" class="form-control" value="" placeholder="Enter the role name" tabindex="1">
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Add" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $("#createRole").on('show.bs.modal', function() {
        $userId = $(event.target).closest('tr').data('id');
        $(this).find("#userId").val($userId);
        console.log($userId)
        setTimeout(function() {
            $('#roleName').focus();
        }, 500);
    });
});
</script>

<div class="modal fade" id="deleteRole" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Delete Role</h4>
            </div>
            <form action="{{urlfor "RoleController.Post"}}" method="POST">
                <div class="modal-body">
                    <input type="hidden" name="userId" id="userId" class="form-control" value="" placeholder="Enter the role name">
                    <input type="hidden" name="roleName" id="roleName" class="form-control" value="" placeholder="Enter the role name">
                    <p>Delete role '<span id="roleNameShown"></span>' ?</p>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Confirm" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $('#deleteRole').on('show.bs.modal', function() {
        $userId = $(event.target).data('id');
        $roleName = $(event.target).data('rolename');
        $(this).find("#userId").val($userId);
        $(this).find("#roleName").val($roleName);
        $(this).find("#roleNameShown").html($roleName);
    });
});
</script>

<div class="modal fade" id="createApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Add App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <input type="hidden" name="userId" id="userId" class="form-control" value="">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="appName" id="appName" class="form-control" value="" placeholder="Enter the app name" tabindex="1">
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Add" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $("#createApp").on('show.bs.modal', function() {
        $userId = $(event.target).closest('tr').data('id');
        $(this).find("#userId").val($userId);
        console.log($userId)
        setTimeout(function() {
            $('#appName').focus();
        }, 500);
    });
});
</script>

<div class="modal fade" id="deleteApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Delete App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <div class="modal-body">
                    <input type="hidden" name="userId" id="userId" class="form-control" value="">
                    <input type="hidden" name="appName" id="appName" class="form-control" value="">
                    <p>Delete app '<span id="appNameShown"></span>' ?</p>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Confirm" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $('#deleteApp').on('show.bs.modal', function() {
        $userId = $(event.target).data('id');
        $appName = $(event.target).data('appname');
        $(this).find("#userId").val($userId);
        $(this).find("#appName").val($appName);
        $(this).find("#appNameShown").html($appName);
    });
});
</script>
