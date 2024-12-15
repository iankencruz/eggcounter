<script>
	import { page } from '$app/stores';
	$effect(() => {
		if (page.status === 401) {
			window.location.href = '/auth/login';
		}
	});
</script>

<h1>{$page.status}: -=-=-> {$page.error.message}</h1>
