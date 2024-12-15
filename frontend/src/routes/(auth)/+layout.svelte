<script>
	export let children;
</script>

<div class="flex min-h-full flex-col justify-center px-6 lg:px-8">

	<div class="sm:mx-auto sm:w-full sm:max-w-sm">
		{@render children()}
	</div>
</div>
