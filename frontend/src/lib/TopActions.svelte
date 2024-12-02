<script lang="ts">
    import { _ } from "svelte-i18n";

    let {
        isEntriesList = $bindable(),
        search = $bindable(),
    }: {
        isEntriesList: boolean;
        search: string;
    } = $props();
</script>

<!-- {$_("input_top_actions")} -->

<div class="absolute top-3 left-2 right-4 flex gap-3">
    <!-- svelte-ignore a11y_autofocus -->
    <input
        class="input input-bordered input-primary input-sm bg-slate-800 flex-1"
        bind:value={search}
        onfocus={() => (isEntriesList = false)}
        id="search"
        type="text"
        placeholder={$_("input_top_actions")}
    />
    <button
        data-tip={$_("search_tip")}
        class="btn btn-ghost btn-sm tooltip tooltip-bottom text-primary flex gap-2"
        onclick={() => (search ? (isEntriesList = !isEntriesList) : false)}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="#7480ff"
            class="w-5"
            viewBox="0 0 16 16"
        >
            <path
                d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001q.044.06.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1 1 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0"
            />
        </svg>

        {$_("search_btn")}
    </button>
    <button
        data-tip={$_("entries_list_tip")}
        class="btn btn-ghost btn-sm tooltip tooltip-left"
        aria-label="Entries list"
        onclick={() => {
            search = "";
            isEntriesList = !isEntriesList;
        }}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="#7480ff"
            class="w-5"
            viewBox="0 0 16 16"
        >
            <path
                fill-rule="evenodd"
                d="M2 2.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5zM3 3H2v1h1z"
            />
            <path
                d="M5 3.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5M5.5 7a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1zm0 4a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1z"
            />
            <path
                fill-rule="evenodd"
                d="M1.5 7a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zM2 7h1v1H2zm0 3.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm1 .5H2v1h1z"
            />
        </svg>
    </button>
</div>
