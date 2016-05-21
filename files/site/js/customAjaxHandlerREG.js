$( "#businessRegisterForm" ).submit(function(event) {
<<<<<<< HEAD
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
	
  var posting = $.post( urlI, { name: nameI, type: typeI, city: cityI, industry: industryI, description: descriptionI, email: emailI, password: passwordI})
  .always(function() {
	
  })
  .done(function() {
	
  })
  .fail(function() {
	
  });
});
=======
  // event.preventDefault();

  var inp = $("input[name='passw']");
  inp.val(md5(inp.val()));
	// nameI = $form.find( "input[name='name']" ).val(),
	// typeI = $form.find( "input[name='type']" ).val(),
	// cityI = $form.find( "input[name='city']" ).val(),
	// industryI = $form.find( "input[name='industry']" ).val(),
	// descriptionI = $form.find( "textarea[name='description']" ).val(),
	// emailI = $form.find( "input[name='email']" ).val(),
	// passwI = $form.find( "input[name='passw']" ).val(),
	// urlI = $form.attr( "action" );

  // var posting = $.post( urlI, { name: nameI, type: typeI, city: cityI, industry: industryI, description: descriptionI, email: emailI, password: passwordI})
  // .always(function() {
	// alert("Finish");
  // })
  // .done(function() {
  //   alert("done");
  // })
  // .fail(function() {
  //   alert("error");
  // });
  return true;
});
>>>>>>> 7639a3c24a275e7e4eb6b4c2ea9c7a5800eb6dab
