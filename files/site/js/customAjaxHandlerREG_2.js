// $( "#businessRegisterForm" ).submit(function(event) {
  // event.preventDefault();

function listEr() {
  var $form = $( "#businessRegisterForm" ),
	nameI = $form.find( "input[name='name']" ).val(),
	typeI = $form.find( "input[name='type']" ).val(),
	cityI = $form.find( "input[name='city']" ).val(),
	industryI = $form.find( "input[name='industry']" ).val(),
	descriptionI = $form.find( "textarea[name='description']" ).val(),
	emailI = $form.find( "input[name='email']" ).val(),
	passwordI = $form.find( "input[name='passw']" ).val(),
	urlI = $form.attr( "action" );
	window.location.replace("/business-reg/check/" + nameI + "/" + typeI + "/" + cityI + "/" + industryI + "/" + descriptionI + "/" + emailI + "/" + passwI);
  // var posting = $.post( urlI, { name: nameI, type: typeI, city: cityI, industry: industryI, description: descriptionI, email: emailI, password: passwordI})
  // .always(function() {
	//
  // })
  // .done(function() {
	//
  // })
  // .fail(function() {
	//
  // });
  return false;
}
