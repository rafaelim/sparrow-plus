<script lang="ts">
	import videoJS from 'video.js';

	type VideoOptions = {
		source: string;
	};
	export let options: VideoOptions;

	let scrollbox: Element;
	$: if (scrollbox) {
		const player = videoJS(scrollbox, {
			controls: true,
			autoplay: true,
			aspectRatio: '16:9',
			html5: {
				nativeTextTracks: false,
				nativeAudioTracks: false,
				hls: {
					overrideNative: true
				}
			},
			sources: [
				{
					src: options.source,
					type: 'application/x-mpegURL'
				}
			]
		});
	}
</script>

<!-- svelte-ignore a11y-media-has-caption -->
<video bind:this={scrollbox} controls={true} class="vjs-matrix video-js" />
<div id="audioTrackControl"></div>
