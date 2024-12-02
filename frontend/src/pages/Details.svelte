<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { copy } from "svelte-copy";
    import BackBtn from "../lib/BackBtn.svelte";
    import ShowPasswordBtn from "../lib/ShowPasswordBtn.svelte";
    import { GetEntryById } from "../../wailsjs/go/main/App";
    import { models } from "../../wailsjs/go/models";
    import EditActions from "../lib/EditActions.svelte";

    let mounted: boolean = $state(false);
    let entry: models.PasswordEntry = $state({} as models.PasswordEntry);
    let show: boolean = $state(false);
    let visible: boolean = $state(false);
    let toast: string = $state("");
    let tmId: number = $state(0);

    let { params }: { params: { id: string } } = $props();

    onMount(() => {
        mounted = true;

        GetEntryById(params.id).then((result) => (entry = result));

        return () => clearTimeout(tmId);
    });

    const onCopy = (element: string) => {
        toast = `${element} copied to clipboard.`;
        visible = true;
        tmId = setTimeout(() => {
            toast = "";
            visible = false;
        }, 2000);
    };
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen relative"
    >
        <div
            class="w-full m-auto flex flex-col gap-2 bg-slate-800 rounded-lg p-2 ml-2 mr-6 relative"
        >
            <h1 class="text-xl font-light m-auto mb-6">Password Details</h1>
            <p class="text-start">{entry.Website}</p>
            <div class="flex gap-2">
                <p class="text-start flex-1">{entry.Username}</p>
                <button
                    data-tip="Copy"
                    class="btn btn-ghost btn-sm tooltip tooltip-top"
                    aria-label="Copy"
                    use:copy={`${entry.Username}`}
                    onclick={() => onCopy("Username")}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="#7480ff"
                        class="w-5"
                        viewBox="0 0 16 16"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M4 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm2-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zM2 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1h1v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1v1z"
                        />
                    </svg>
                </button>
            </div>
            <div class="flex gap-2">
                <div class="relative flex-1">
                    <input
                        class="input input-bordered input-primary input-sm bg-slate-800 w-full"
                        bind:value={entry.Password}
                        id="password"
                        type={show ? "text" : "password"}
                        disabled
                    />
                    <ShowPasswordBtn bind:showPass={show} />
                </div>
                <button
                    data-tip="Copy"
                    class="btn btn-ghost btn-sm tooltip tooltip-top"
                    aria-label="Copy"
                    use:copy={`${entry.Password}`}
                    onclick={() => onCopy("Password")}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="#7480ff"
                        class="w-5"
                        viewBox="0 0 16 16"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M4 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm2-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zM2 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1h1v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1v1z"
                        />
                    </svg>
                </button>
            </div>
        </div>
        {#if visible}
            <p
                out:fade={{ duration: 800 }}
                class="absolute bottom-10 text-xs font-bold text-green-400 inset-x-0 justify-center"
            >
                {toast}
            </p>
        {/if}

        <BackBtn />
        <EditActions {entry} />
    </div>
{/if}
