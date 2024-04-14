import { writable, derived } from 'svelte/store';

type Episode = {
	episodeId: string;
	episodeNumber: string;
	name: string;
	season: string;
	path: string;
};

export const apiData = writable<Episode[]>([]);

export const episodes = derived(apiData, ($apiData) => {
	return $apiData;
});
