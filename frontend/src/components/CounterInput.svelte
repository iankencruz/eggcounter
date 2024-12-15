<script lang="ts">
	let { class: className, count, saveFunction } = $props();

	let inputCount = $state(0); // Local count for the user's current input
	let saving = $state(false);

	const increase = () => inputCount++;
	const decrease = () => {
		if (inputCount > 0) inputCount--;
	};

	const save = async () => {
		if (saving) return; // Prevent multiple saves
		saving = true;
		try {
			await saveFunction(inputCount); // Call the parent-provided function
			inputCount = 0;
		} catch (err) {
			console.error('Failed to save:', err);
		} finally {
			saving = false;
		}
	};
</script>

<div class={className}>
	<div class="flex w-full space-x-4">
		<!-- Decrease button -->
		<button
			disabled={count <= 0 || saving}
			onclick={decrease}
			class="shadow-xs rounded-full bg-indigo-600 px-4 py-1.5 text-white hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
		>
			-
		</button>

		<!-- Input field for count -->
		<input
			type="number"
			bind:value={inputCount}
			class="w-16 rounded border border-gray-300 text-center focus:outline-none focus:ring focus:ring-indigo-300"
			min="0"
		/>

		<!-- Increase button -->
		<button
			onclick={increase}
			disabled={saving}
			class="shadow-xs rounded-full bg-indigo-600 px-4 py-1.5 text-white hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
		>
			+
		</button>
	</div>

	<!-- Save button -->
	<button
		onclick={save}
		disabled={saving}
		class="ml-auto rounded bg-indigo-500 px-4 py-2 font-semibold text-white hover:bg-indigo-600 disabled:opacity-50"
	>
		{#if saving}
			Saving...
		{/if}
		{#if !saving}
			Save
		{/if}
	</button>
</div>

<style>
	input::-webkit-inner-spin-button {
		appearance: none;
		margin: 0;
	}

	input[type='number'] {
		-moz-appearance: textfield;
		appearance: textfield;
	}

	.disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}
</style>
