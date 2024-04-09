<script lang="ts">
	import 'video.js/dist/video-js.css';
	import videoJS from 'video.js';

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
					src: 'http://localhost:3000/api/stream/<filename>/master',
					type: 'application/x-mpegURL'
				}
			]
		});
	}
</script>

<!-- svelte-ignore a11y-media-has-caption -->
<video bind:this={scrollbox} controls={true} class="vjs-matrix video-js" />
<div id="audioTrackControl"></div>
