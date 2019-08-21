<header>
        <nav class="navbar navbar-inverse">
            <div class="container-fluid">
              <div class="navbar-header">
                <a class="navbar-brand" href="/">FluxWeb</a>
              </div>
              <ul class="nav navbar-nav">
                <li class="active"><a href="/">Home</a></li>
                {{if lt 0 (len .namespaces)}}
                {{range $ns := .namespaces}}
                  <li><a href="/workload/{{$ns}}">{{$ns}}</a></li>
                {{ end }}{{ end }}
              </ul>
              <form id="form" class="navbar-form navbar-left"> 
                <div class="form-group">
                  <input id="namespace_input" type="text" class="form-control" name="search" placeholder="namespace">
                </div>
                <button id="submit_btn" type="submit" class="btn btn-default">Submit</button>
            </form> 
            </div>
        </nav>
</header>
