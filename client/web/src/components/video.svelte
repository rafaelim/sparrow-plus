<script lang="ts">
    import 'video.js/dist/video-js.css';
	import videoJS from 'video.js';

    type VideoOptions = {
        source: string
    }
    export let options: VideoOptions
    
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
<video {...options} bind:this={scrollbox} controls={true} class="vjs-matrix video-js" />