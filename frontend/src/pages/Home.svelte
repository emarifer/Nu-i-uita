<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { _ } from "svelte-i18n";
    import BottomActions from "../lib/BottomActions.svelte";
    import TopActions from "../lib/TopActions.svelte";
    import { GetPasswordCount } from "../../wailsjs/go/main/App";
    import EntriesList from "../lib/EntriesList.svelte";

    let mounted: boolean = false,
        count: number = 0,
        showList: boolean = false,
        searchTerms: string = "";

    onMount(() => {
        mounted = true;
        GetPasswordCount().then((result) => (count = result));
    });
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen relative"
    >
        <TopActions bind:isEntriesList={showList} bind:search={searchTerms} />

        {#if !showList}
            <h1 class="text-lg font-light m-auto">
                {count}&nbsp;{$_("home_title")}
            </h1>
        {:else}
            <EntriesList bind:listCounter={count} bind:search={searchTerms} />
        {/if}

        <BottomActions />
    </div>
{/if}
