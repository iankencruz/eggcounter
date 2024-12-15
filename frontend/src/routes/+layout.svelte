<script lang="ts">
	import '../app.css';

	// Access children routes
	let { children } = $props();
</script>

<div class="h-full bg-gray-50">
	<!-- Main Content Area -->
	<main class="min-h-svh py-4">
		<div class="mx-auto">
			{@render children()}
		</div>
	</main>
</div>

<style>
	.h-full {
		height: 100%;
	}
</style>
