<script lang="ts">
    import { models } from "../../wailsjs/go/models";
    import {
        GetAllEntries,
        GetPasswordCount,
        DeleteEntry,
    } from "../../wailsjs/go/main/App";
    import { onDestroy, onMount } from "svelte";
    import { _ } from "svelte-i18n";
    import { push } from "svelte-spa-router";
    import { showSuccess, showWarning } from "./popups/popups";

    let {
        listCounter = $bindable(),
        search = $bindable(),
    }: {
        listCounter: number;
        search: string;
    } = $props();

    let entries: models.PasswordEntry[] = $state([]);

    onMount(() => {
        GetAllEntries().then((result) => {
            // console.log("SEARCH:", search);
            if (search) {
                const find = search.toLowerCase();
                entries = result.filter(
                    (item) =>
                        item.Username.toLowerCase().includes(find) ||
                        item.Website.toLowerCase().includes(find),
                );
            } else {
                entries = result;
            }
        });
    });

    onDestroy(() => (search = ""));

    const showAlert = (website: string, id: string) => {
        const data: string[] = [
            `${$_("alert_deleting_password")} "${website}."`,
            `${$_("alert_confirm_deleting")}`,
        ];
        showWarning(data).then((result) => {
            if (result.value) {
                DeleteEntry(id).then(() => deleteItem(id));
                showSuccess($_("deletion_confirm_msg"));
            }
        });
    };

    const deleteItem = (id: string): void => {
        let itemIdx = entries.findIndex((x) => x.Id === id);
        entries.splice(itemIdx, 1);
        entries = entries;
        GetPasswordCount().then((result) => (listCounter = result));
    };
</script>

{#if search}
    <h1 class="font-light absolute top-14 right-0 left-0">
        {$_("search_result")}
    </h1>
{/if}
{#if entries.length}
    <ul
        class="w-full m-auto max-h-40 overflow-y-auto flex flex-col gap-2 bg-slate-800 rounded-lg p-2 ml-2 mr-4 overflow-x-hidden"
    >
        {#each entries as entry (entry.Id)}
            {@render item(entry)}
        {/each}
    </ul>
{:else}
    <h1 class="text-lg font-light m-auto">{$_("no_password_entries")}</h1>
{/if}

{#snippet item(entry: models.PasswordEntry)}
    <li
        data-tip={entry.Website}
        class="flex justify-between items-center bg-slate-700 rounded-md hover:bg-slate-600 w-full"
    >
        <p
            data-tip={entry.Website}
            class="flex items-center justify-start max-w-72 text-ellipsis overflow-x-clip whitespace-nowrap text-sm pl-2 tooltip tooltip-top"
        >
            {entry.Website}
        </p>
        <div class="flex gap-1 w-fit mr-2">
            <button
                data-tip={$_("details_btn_tip")}
                class="btn btn-ghost btn-sm tooltip tooltip-top"
                aria-label="Details"
                onclick={() => push(`/details/${entry.Id}`)}
            >
                <svg
                    fill="currentColor"
                    class="w-4"
                    viewBox="0 0 24 24"
                    version="1.1"
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                >
                    <g
                        id="web-app"
                        stroke="none"
                        stroke-width="1"
                        fill-rule="evenodd"
                    >
                        <g id="search-plus">
                            <path
                                d="M16.3250784,14.8989201 L21.704633,20.2784747 C22.0984557,20.6722974 22.0984557,21.3108102 21.704633,21.704633 C21.3108102,22.0984557 20.6722974,22.0984557 20.2784747,21.704633 L14.8989201,16.3250784 C13.5453296,17.3749907 11.8456542,18 10,18 C5.581722,18 2,14.418278 2,10 C2,5.581722 5.581722,2 10,2 C14.418278,2 18,5.581722 18,10 C18,11.8456542 17.3749907,13.5453296 16.3250784,14.8989201 Z M10,16 C13.3137085,16 16,13.3137085 16,10 C16,6.6862915 13.3137085,4 10,4 C6.6862915,4 4,6.6862915 4,10 C4,13.3137085 6.6862915,16 10,16 Z M13,11 L11,11 L11,13 C11,13.5522847 10.5522847,14 10,14 C9.44771525,14 9,13.5522847 9,13 L9,11 L7,11 C6.44771525,11 6,10.5522847 6,10 C6,9.44771525 6.44771525,9 7,9 L9,9 L9,7 C9,6.44771525 9.44771525,6 10,6 C10.5522847,6 11,6.44771525 11,7 L11,9 L13,9 C13.5522847,9 14,9.44771525 14,10 C14,10.5522847 13.5522847,11 13,11 Z"
                                id="Shape"
                            >
                            </path>
                        </g>
                    </g>
                </svg>
            </button>
            <button
                data-tip={$_("delete_btn_tip")}
                class="btn btn-ghost btn-sm tooltip tooltip-top"
                aria-label="Delete Item"
                onclick={() => showAlert(entry.Website, entry.Id)}
            >
                <svg
                    fill="currentColor"
                    class="w-4"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 16 16"
                >
                    <path
                        d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5M8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5m3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0"
                    />
                </svg>
            </button>
        </div>
    </li>
{/snippet}

<!-- SVELTE. DELETE OBJECT FROM ARRAY AND RE-RENDER VIEW:
    https://svelte.dev/playground/a8dc7bb4c7744e36826fb4b1b3d12195?version=3.23.2
 -->
