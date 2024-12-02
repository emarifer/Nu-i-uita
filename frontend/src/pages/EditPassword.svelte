<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import BackBtn from "../lib/BackBtn.svelte";
    import ShowPasswordBtn from "../lib/ShowPasswordBtn.svelte";
    import { GetEntryById, UpdateEntry } from "../../wailsjs/go/main/App";
    import { models } from "../../wailsjs/go/models";
    import { push } from "svelte-spa-router";

    let mounted: boolean = $state(false);
    let entry: models.PasswordEntry = $state.snapshot(
        {} as models.PasswordEntry,
    );
    let website: string = $state("");
    let username: string = $state("");
    let password: string = $state("");
    let show: boolean = $state(false);
    let visible: boolean = $state(false);
    let toast: string = $state("");
    let tmId1: number = $state(0);
    let tmId2: number = $state(0);

    let { params }: { params: { id: string } } = $props();

    onMount(() => {
        mounted = true;

        GetEntryById(params.id).then((result) => {
            entry = result;
            website = result.Website;
            username = result.Username;
            password = result.Password;
        });

        return () => {
            clearTimeout(tmId1);
            clearTimeout(tmId2);
            // console.log("Destroying the component…");
        };
    });

    const onUpdate = () => {
        if (
            website.trim() === "" ||
            username.trim() === "" ||
            password.trim() === ""
        ) {
            toast = "None of the fields can be empty !!";
            visible = true;

            tmId1 = setTimeout(() => {
                toast = "";
                visible = false;
            }, 2000);
            website = entry.Website;
            username = entry.Username;
            password = entry.Password;
            return;
        } else if (password.trim().length < 6) {
            toast = "Password must be at least 6 characters long !!";
            visible = true;

            tmId2 = setTimeout(() => {
                toast = "";
                visible = false;
            }, 2000);
            website = entry.Website;
            username = entry.Username;
            password = entry.Password;
            return;
        }

        entry.Website = website;
        entry.Username = username;
        entry.Password = password;
        // console.log("Entry Password", entry);

        UpdateEntry(entry).then((result) => {
            result ? push("/home") : false;
        });
    };
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen relative"
    >
        <div class="flex flex-col gap-3 w-full m-auto pl-6 pr-10">
            <label for="website" class="text-lg font-light mx-auto">
                Edit Entry Password
            </label>
            <input
                class="input input-bordered input-primary input-sm w-full bg-slate-800"
                autocomplete="off"
                bind:value={website}
                id="website"
                type="text"
                placeholder="Website"
            />
            <input
                class="input input-bordered input-primary input-sm w-full bg-slate-800"
                autocomplete="off"
                bind:value={username}
                id="username"
                type="text"
                placeholder="Username"
            />
            <div class="relative">
                <input
                    class="input input-bordered input-primary input-sm w-full bg-slate-800"
                    autocomplete="off"
                    bind:value={password}
                    id="website"
                    type={show ? "text" : "password"}
                    placeholder="Password"
                />
                <ShowPasswordBtn bind:showPass={show} />
            </div>
            <div class="flex relative">
                <button
                    class="btn btn-sm btn-outline btn-primary w-full mb-8"
                    onclick={onUpdate}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        class="w-5"
                        viewBox="0 0 16 16"
                    >
                        <path
                            d="M0 1.5A1.5 1.5 0 0 1 1.5 0H3v5.5A1.5 1.5 0 0 0 4.5 7h7A1.5 1.5 0 0 0 13 5.5V0h.086a1.5 1.5 0 0 1 1.06.44l1.415 1.414A1.5 1.5 0 0 1 16 2.914V14.5a1.5 1.5 0 0 1-1.5 1.5H14v-5.5A1.5 1.5 0 0 0 12.5 9h-9A1.5 1.5 0 0 0 2 10.5V16h-.5A1.5 1.5 0 0 1 0 14.5z"
                        />
                        <path
                            d="M3 16h10v-5.5a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5zm9-16H4v5.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5zM9 1h2v4H9z"
                        />
                    </svg>
                    Update Entry
                </button>
                {#if visible}
                    <p
                        out:fade={{ duration: 800 }}
                        class="absolute top-9 text-xs font-bold text-secondary inset-x-0 justify-center"
                    >
                        {toast}
                    </p>
                {/if}
            </div>
        </div>

        <BackBtn />
    </div>
{/if}
