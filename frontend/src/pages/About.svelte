<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { _ } from "svelte-i18n";
    import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";
    import { GetRepositoryTag } from "../../wailsjs/go/main/App";
    import BackBtn from "../lib/BackBtn.svelte";

    let mounted: boolean = false;
    let version: string = "";

    onMount(() => {
        mounted = true;
        GetRepositoryTag().then((result) => (version = result));
    });
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen relative"
    >
        <div class="flex flex-col gap-4 m-auto">
            <h1 class="text-lg font-light mb-8">{$_("about")}</h1>
            <h3 class="text-xs font-thin text-amber-600">
                {$_("app_description")}
            </h3>
            <p class="text-xs text-slate-600 font-medium">{$_("more_info")}</p>
            <a
                onclick={() =>
                    BrowserOpenURL(
                        "https://github.com/emarifer/go-wails-svelte-desktop-todoapp",
                    )}
                class="text-xs font-medium hover:text-sky-500 ease-in duration-300 -m-4"
                href="http://"
                target="_blank"
                rel="noopener noreferrer"
            >
                https://github.com/emarifer/go-wails-svelte-desktop-todoapp
            </a>
        </div>

        <span class="text-xs font-bold text-sky-600 absolute bottom-4 right-4">
            {$_("version")}&nbsp;{version}
        </span>

        <BackBtn />
    </div>
{/if}
