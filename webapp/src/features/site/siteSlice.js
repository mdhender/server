import {createSlice} from '@reduxjs/toolkit';

export const siteSlice = createSlice({
	name: 'site',
	initialState: {
		copyright: {
			author: 'Michael D Henderson',
			year: 2020,
		},
		title: 'Gas Giant Battles',
		slug: 'The client for the server...',
		notice: {
			title: 'License',
			text: 'Gas Giant Battles is free software, licensed under the AGPL v3.',
		},
		menus: {
			nav: [
				{
					id: 1, link: "/", label: "Main pages", class: "nav", children: [
						{id: 1.1, link: "/blog", label: "Blog"},
						{id: 1.2, link: "/forums", label: "Forums"},
						{id: 1.3, link: "/about", label: "About"},
					]
				},
				{
					id: 2, link: "/", label: "Games", class: "nav", children: [
						{id: 2.1, link: "/rules", label: "Rules"},
						{id: 2.2, link: "/tournaments", label: "Tournaments"},
						{id: 2.3, link: "/archives", label: "Archives"},
					]
				},
				{
					id: 3, link: "/", label: "Extras", class: "nav", children: [
						{id: 3.1, link: "/extras/music-archive", label: "Music archive"},
						{id: 3.2, link: "/extras/photo-gallery", label: "Photo gallery"},
						{id: 3.3, link: "/extras/poems-and-lyrics", label: "Poems and lyrics"},
					]
				},
				{
					id: 4, link: "/", label: "Community", class: "nav", children: [
						{id: 4.1, link: "/guestbook", label: "Guestbook"},
						{id: 4.2, link: "/members", label: "Members"},
						{id: 4.3, link: "/links", label: "Link collection"},
					]
				},
				{
					id: 5, link: "/status", label: "System Status", class: "introduction", children: [
						{id: 5.1, label: "Game 103 now accepting new players."},
						{id: 5.2, label: "Usagi's attack on 05.07.13 repelled by Chizu."},
					]
				},
			],
			side: {
				left: [
					{
						id: 1, link: "/my/games", label: "My Games", children: [
							{id: 1.1, link: "/my/games/100", label: "Game 100"},
							{id: 1.2, link: "/my/games/102", label: "Game 102"},
						]
					},
					{id: 2, link: "/games/active", label: "Active Games"},
					{
						id: 3, link: "/admin", label: "Administration", children: [
							{id: 3.1, link: "/admin/create-game", label: "Create Game"},
							{id: 3.2, link: "/admin/upgrade-server", label: "Upgrade Server"},
						]
					},
				],
				right: [
					{
						id: 1, link: "/my/games/100", label: "Overview", children: [
							{id: 1.1, link: "/my/games/100/reports", label: "Reports"},
							{id: 1.2, link: "/my/games/100/systems", label: "Systems"},
							{id: 1.3, link: "/my/games/100/diplomacy", label: "Diplomacy"},
						]
					},
					{id: 2, link: "/my/games/100/orders", label: "Orders"},
					{
						id: 3, link: "/my/games/100/units", label: "Units", children: [
							{id: 3.1, link: "/my/games/100/units/colonies", label: "Colonies"},
							{id: 3.2, link: "/my/games/100/units/mines", label: "Mines"},
							{id: 3.3, link: "/my/games/100/units/factories", label: "Factories"},
							{id: 3.4, link: "/my/games/100/units/ships", label: "Ships"},
							{id: 3.5, link: "/my/games/100/units/population", label: "Population"},
						]
					},
					{id: 4, link: "/my/games/100/to-do", label: "To Do"},
					{id: 5, link: "/my/games/100/history", label: "History"},
				],
				samples: [
					{id: 1, link: "/sample-link-1", label: "Sample link 1"},
					{id: 2, link: "/sample-link-2", label: "Sample link 2"},
					{id: 3, link: "/sample-link-3", label: "Sample link 3"},
					{id: 4, link: "/sample-link-4", label: "Sample link 4"},
				],
			},
		},
	},
});

export const selectCopyright = state => state.site.copyright;
export const selectSlug = state => state.site.slug;
export const selectTitle = state => state.site.title;

export default siteSlice.reducer;
