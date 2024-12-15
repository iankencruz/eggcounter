 <script lang="ts">
	import '../../app.css';
	import { goto } from '$app/navigation';
	import SideNav from '../../components/SideNav.svelte';

/*
	

	export const load = async ({ fetch }) => {
		const res = await fetch('/api/auth/status');

		console.log(res.status);
		if (res.status === 401) {
			// Redirect on unauthorized
			return { status: 302, redirect: '/login' };
		}

		const data = await res.json();
		return { props: { user: data.user || null } };
	};
*/
	// Side effects for document title and redirect
	$effect(() => {
		// Set the document title for this layout
		document.title = 'Dashboard';

		// Redirect to /dashboard if at /
		if (window.location.pathname === '/') {
			goto('/login'); // Use SvelteKit's client-side navigation
		}
	});
	



// Access children routes
	let { children } = $props();

</script>
<div class="h-full bg-gray-50">

	<SideNav />


	<main class="min-h-svh py-10 lg:pl-72">
		<div class="px-4 sm:px-6 lg:px-8">
			{@render children()}
		</div>
	</main>
</div>

<style>
	.h-full {
		height: 100%;
	}
</style> 
