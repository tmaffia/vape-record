$(document).ready(function() {
	$('.record-vape').click(function(){
		$.get("api/vape", function() {
			location.reload();
			});
	});

	$('.delete-last-record').click(function() {
		$.get("api/delete-vape", function() {
			location.reload();
			});
	});
});
