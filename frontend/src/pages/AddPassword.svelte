<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import BackBtn from "../lib/BackBtn.svelte";
    import ShowPasswordBtn from "../lib/ShowPasswordBtn.svelte";
    // import Swal from "sweetalert2";
    // import warning from "../assets/images/warning.png";
    import { AddPasswordEntry } from "../../wailsjs/go/main/App";
    import { push } from "svelte-spa-router";

    let mounted: boolean = false,
        inputRef: HTMLInputElement | null = null,
        website: string = "",
        username: string = "",
        newPassword: string = "",
        show = false,
        visible = false,
        toast = "",
        tmId1 = 0,
        tmId2 = 0;

    /* const showAlert = () =>
        Swal.fire({
            title: "Invalid input",
            width: 275,
            icon: "warning",
            iconHtml: `<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" viewBox="0 0 16 16">
  <path d="M11.46.146A.5.5 0 0 0 11.107 0H4.893a.5.5 0 0 0-.353.146L.146 4.54A.5.5 0 0 0 0 4.893v6.214a.5.5 0 0 0 .146.353l4.394 4.394a.5.5 0 0 0 .353.146h6.214a.5.5 0 0 0 .353-.146l4.394-4.394a.5.5 0 0 0 .146-.353V4.893a.5.5 0 0 0-.146-.353zM8 4c.535 0 .954.462.9.995l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995A.905.905 0 0 1 8 4m.002 6a1 1 0 1 1 0 2 1 1 0 0 1 0-2"/>
</svg>`,
            background: "#1D232A",
            color: "#A6ADBA",
            confirmButtonColor: "#3085d6",
            customClass: {
                title: "alertTitle",
                confirmButton: "alertConfirm",
                icon: "alertIcon",
            },
        }); */

    onMount(() => {
        mounted = true;

        return () => {
            clearTimeout(tmId1);
            clearTimeout(tmId2);
            // console.log("Destroying the component…");
        };
    });

    const onSave = () => {
        if (
            website.trim() === "" ||
            username.trim() === "" ||
            newPassword.trim() === ""
        ) {
            toast = "None of the fields can be empty !!";
            visible = true;

            tmId1 = setTimeout(() => {
                toast = "";
                visible = false;
            }, 2000);
            website = "";
            username = "";
            newPassword = "";
            inputRef?.focus();
            return;
        } else if (newPassword.trim().length < 6) {
            toast = "Password must be at least 6 characters long !!";
            visible = true;

            tmId2 = setTimeout(() => {
                toast = "";
                visible = false;
            }, 2000);
            website = "";
            username = "";
            newPassword = "";
            inputRef?.focus();
            return;
        }

        AddPasswordEntry(website, username, newPassword).then((result) => {
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
                Add New Password
            </label>
            <!-- svelte-ignore a11y_autofocus -->
            <input
                class="input input-bordered input-primary input-sm w-full bg-slate-800"
                autocomplete="off"
                bind:this={inputRef}
                bind:value={website}
                id="website"
                type="text"
                placeholder="Website"
                autofocus
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
                    bind:value={newPassword}
                    id="website"
                    type={show ? "text" : "password"}
                    placeholder="Password"
                />
                <ShowPasswordBtn bind:showPass={show} />
            </div>
            <div class="flex relative">
                <button
                    class="btn btn-sm btn-outline btn-primary w-full mb-8"
                    onclick={onSave}
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
                    Create Entry
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

<!-- CUSTOMIZATION BY SWEETALERT2:
    https://sweetalert2.github.io/recipe-gallery/custom-icon.html
 -->
