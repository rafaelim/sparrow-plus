export type CarouselOptions = {
	titleKey?: string;
	idKey?: string;
	nextRoute?: string | ((row: Record<string, unknown>) => string);
	getParams?: (row: Record<string, unknown>) => string;
};
