<!DOCTYPE html>
<html>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>
    <link href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js"></script>
    <script src='/static/js/main.js'></script>
    <link rel="stylesheet" type="text/css" href="/static/css/main.css" media="screen" />
    <!------ Include the above in your HEAD tag ---------->
<body> 
  
  <section>
      {{ template "header.tpl" . }}
  </section>

  <form id="filter_form" class="navbar-form"> 
    <div class="form-group">
      <input id="filter_input" type="text" name="filter" placeholder="filter workloads">
    </div>
    <!--<button id="filter_btn" type="submit" class="btn btn-default">Filter</button>-->
  </form> 

  <section>
    {{ template "workloads.tpl" . }}
  </section>

</body>  
    
</html>
