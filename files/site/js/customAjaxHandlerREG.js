$( "#businessRegisterForm" ).submit(function(event) {
  event.preventDefault();
  
  var $form = $( this ),
	nameI = $form.find( "input[name='name']" ).val(),
	typeI = $form.find( "input[name='type']" ).val(),
	cityI = $form.find( "input[name='city']" ).val(),
	industryI = $form.find( "input[name='industry']" ).val(),
	descriptionI = $form.find( "textarea[name='description']" ).val(),
	emailI = $form.find( "input[name='email']" ).val(),
	passwordI = $form.find( "input[name='passw']" ).val(),
	urlI = $form.attr( "action" );
	
  var posting = $.post( urlI, { name: nameI, type: typeI, city: cityI, industry: industryI, description: descriptionI, email: emailI, password: passwordI});
  posting.done(function() {
    window.location = "profile_business.html";
  })
  posting.fail(function() {
	window.location = "profile_business.html";
  });
});