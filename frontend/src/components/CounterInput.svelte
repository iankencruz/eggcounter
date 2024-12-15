<script lang="ts">
	let { class: className, count } = $props();

	const increase = () => count++;
	const decrease = () => {
		if (count > 0) count--;
	};

	const save = () => {
		alert(`Saved value: ${count}`);
	};
</script>

<div class={className}>
	<div class="flex items-center justify-between gap-4 rounded border border-gray-300 p-2">
		<!-- Minus Button -->
		<button
			onclick={decrease}
			class="rounded bg-gray-200 px-3 py-2 font-semibold text-gray-800 hover:bg-gray-300"
		>
			-
		</button>

		<!-- Number Input -->
		<input
			type="number"
			bind:value={count}
			class="w-16 rounded border border-gray-300 text-center focus:outline-none focus:ring focus:ring-indigo-300"
			min="0"
		/>

		<!-- Plus Button -->
		<button
			onclick={increase}
			class="rounded bg-gray-200 px-3 py-2 font-semibold text-gray-800 hover:bg-gray-300"
		>
			+
		</button>

		<!-- Save Button -->
		<button
			onclick={save}
			class="ml-auto rounded bg-indigo-500 px-4 py-2 font-semibold text-white hover:bg-indigo-600"
		>
			Save
		</button>
	</div>
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
</style>
