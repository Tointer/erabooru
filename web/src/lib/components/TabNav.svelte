<script lang="ts">
	/**
	 * Page currently shown so we can highlight the active tab
	 */
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import { PAGE_SIZE } from '$lib/constants';

	let q: string = $state('');
	let active: 'media' | 'upload' | 'settings' = $props();

	onMount(() => {
		q = get(page).url.searchParams.get('q') ?? '';
	});

	function search(event: Event) {
		event.preventDefault();
		goto(`/?q=${encodeURIComponent(q)}&page=1&page_size=${PAGE_SIZE}`);
	}
</script>

<div class="mb-4 border-b">
	<nav class="flex items-center space-x-4">
		<a
			href="/"
			class="-mb-px border-b-2 px-3 py-2"
			class:!border-blue-500={active === 'media'}
			class:!text-blue-500={active === 'media'}
			class:border-transparent={active !== 'media'}
			class:text-gray-500={active !== 'media'}
		>
			Media
		</a>
		<a
			href="/upload"
			class="-mb-px border-b-2 px-3 py-2"
			class:!border-blue-500={active === 'upload'}
			class:!text-blue-500={active === 'upload'}
			class:border-transparent={active !== 'upload'}
			class:text-gray-500={active !== 'upload'}
		>
			Upload
		</a>
		<a
			href="/settings"
			class="-mb-px border-b-2 px-3 py-2"
			class:!border-blue-500={active === 'settings'}
			class:!text-blue-500={active === 'settings'}
			class:border-transparent={active !== 'settings'}
			class:text-gray-500={active !== 'settings'}
		>
			Settings
		</a>
		<form onsubmit={search} class="mx-auto">
			<input
				type="text"
				name="search"
				placeholder="Search"
				bind:value={q}
				class="rounded border px-2 py-1"
			/>
		</form>
	</nav>
</div>
