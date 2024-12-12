<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { _, locale } from "svelte-i18n";
    import {
        CheckMasterPassword,
        GetLanguage,
        GetMasterPassword,
        SaveMasterPassword,
    } from "../../wailsjs/go/main/App";
    import { push } from "svelte-spa-router";
    import ShowPasswordBtn from "../lib/ShowPasswordBtn.svelte";
    import { EventsEmit } from "../../wailsjs/runtime/runtime";

    let mounted = false,
        inputRef: HTMLInputElement | null = null,
        isLogin = false,
        show = false,
        newPassword = "",
        visible = false,
        toast = "",
        tmId1 = 0,
        tmId2 = 0;

    onMount(() => {
        mounted = true;
        GetMasterPassword().then((result) => {
            isLogin = result;
            // console.log("Master password exists in DB:", isLogin);
        });

        GetLanguage().then((result) => {
            locale.set(result);
            EventsEmit(
                "change_titles",
                $_("app_title"),
                $_("select_directory"),
                $_("select_file"),
            );
        });

        return () => {
            clearTimeout(tmId1);
            clearTimeout(tmId2);
            // console.log("Destroying the component…");
        };
    });

    const onLogin = () => {
        if (newPassword.length < 6 || !isAscii(newPassword)) {
            toast = $_("password_too_short_or_non_ascii_chars");
            visible = true;
            tmId1 = setTimeout(() => {
                toast = "";
                visible = false;
            }, 2000);
            newPassword = "";
            inputRef?.focus();
            return;
        } else if (!isLogin) {
            SaveMasterPassword(newPassword).then((result) => {
                // console.log("PASSWORD_ID:", result);
                result ? push("/home") : false;
            });
            return;
        }

        CheckMasterPassword(newPassword).then((result) => {
            if (result) {
                push("/home");
            } else {
                toast = $_("wrong_password");
                visible = true;
                tmId2 = setTimeout(() => {
                    toast = "";
                    visible = false;
                }, 2000);
                newPassword = "";
                inputRef?.focus();
            }
        });
    };

    const isAscii = (str: string): boolean => /^[\x00-\x7F]+$/.test(str);
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen"
    >
        <div class="flex flex-col gap-2 w-fit h-fit m-auto">
            <div class="flex flex-col gap-4">
                <label for="newPassword" class="text-lg font-light">
                    {!isLogin ? $_("generate") : $_("insert")}
                    {$_("master_password")}
                </label>
                <!-- svelte-ignore a11y_autofocus -->
                <div class="relative">
                    <input
                        class="input input-bordered input-primary input-sm bg-slate-800 w-full"
                        autocomplete="off"
                        bind:this={inputRef}
                        bind:value={newPassword}
                        id="newPassword"
                        type={show ? "text" : "password"}
                        autofocus
                    />
                    <ShowPasswordBtn bind:showPass={show} />
                </div>
            </div>

            <div class="flex relative">
                <button
                    class="btn btn-outline btn-primary btn-sm w-full"
                    onclick={onLogin}
                >
                    {$_("login")}
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
    </div>
{/if}

<!-- 
    Toggle Password:
    https://svelte.dev/playground/3632c24a5c094528b961dc2c7ed119a4?version=3.32.3
    https://preline.co/docs/toggle-password.html

    Daisyui colors pallette & OKLCH Color Picker & Converter:
    https://unpkg.com/browse/daisyui@4.12.2/dist/themes.css
    https://oklch.com/
    https://github.com/saadeghi/daisyui/discussions/3096

    Cleanup timeout:
    https://stackoverflow.com/questions/64862161/svelte-store-function-update

    CHECK ASCII CHARACTERS:
    https://quickref.me/check-if-a-string-contains-only-ascii-characters.html
 -->
