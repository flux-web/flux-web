$(document).ready(function(){
    $('[data-container="body"]').tooltip({ container: 'body' })  
});
$("#submit_btn").on("click",function(){
  //TODO: make header div independent 
  window.location.href = "/workload/" + $("#namespace_input").val();
  //$.get("/workload/" + $("#namespace_input").val(), function(data) {
  //  $('#main').html(data);
  //});
});

var typingTimer;
var doneTypingInterval = 1000;
$('#filter_input').keyup(function(){
    clearTimeout(typingTimer);
    if ($('#filter_input').val()) {
        typingTimer = setTimeout(doneTyping, doneTypingInterval);
    }
});

function doneTyping () {
  //TODO: make workloads div independent 
  //$.get("/?filter=" + $("#filter_input").val(), function(data) {
  //  $('#main').html(data);
  //});
  window.location.href = window.location.pathname + "?filter=" + $("#filter_input").val();
}
$("#form").submit(function(event) {
    event.preventDefault();
});