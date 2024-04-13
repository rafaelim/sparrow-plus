<script lang="ts">
	import type { CarouselOptions } from "$lib/carousel";


	export let options: CarouselOptions = {rows: []}

	const getNextRoute = (row: Record<string, string>) => {
		const key = options.idKey ?? ''
		const rowId = row[key]

		return `${options.nextRoute}/${rowId}`
	}

	const getTitle = (row: Record<string, string>) => {
		const key = options.titleKey ?? ''
		return row[key]
	}
</script>

<div class="container"  {...options}>
	{#if options}
		{#each options.rows as row}
			<a href={getNextRoute(row)} class="box">
				<img src="https://assets.nflxext.com/us/boxshots/tv_sdp_s/70143813.jpg" alt="img" />
				<div class="overlay">
					<h3 class="title">{getTitle(row)}</h3>
				</div>
			</a>
		{/each}
	{/if}
</div>

<style lang="scss">
	.container {
		display: flex;
		align-items: center;
		overflow-y: hidden;
		overflow-x: scroll;
		min-height: 215px;
		transition: 500ms;
		scroll-behavior: smooth;
	}
	.box {
		min-width: 250px;
		height: 129px;
		position: relative;
		cursor: pointer;
		transition: transform 0.5s ease;
	}

	.box img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.box:hover {
		transform: scale(1.5);
		z-index: 2;
	}

	.overlay {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: rgba(0, 0, 0, 0.7);
		padding: 5px 10px;
		transition: all 0.5s ease;
		opacity: 0;
		pointer-events: none;
	}

	.box:hover .overlay {
		opacity: 1;
		pointer-events: all;
	}

	.overlay .title {
		color: #e9e9e9;
		font-size: 12px;
		cursor: pointer;
	}

	.overlay .title:hover {
		color: #fff;
	}
</style>
