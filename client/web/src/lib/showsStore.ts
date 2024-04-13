import { writable, derived } from 'svelte/store';

type Show = {
	movieId: string;
	name: string;
	year: string;
	path: string;
};

export const apiData = writable<Show[]>([]);

export const shows = derived(apiData, ($apiData) => {
	return $apiData;
});
