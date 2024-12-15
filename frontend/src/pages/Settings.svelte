<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import BackBtn from "../lib/BackBtn.svelte";
    import Language from "../lib/Language.svelte";
    import { EventsEmit, EventsOn } from "../../wailsjs/runtime/runtime";
    import { Drop, GetLanguage } from "../../wailsjs/go/main/App";
    import { _, locale } from "svelte-i18n";
    import { push } from "svelte-spa-router";
    import {
        showError,
        showQuestion,
        showSuccess,
        showWarning,
    } from "../lib/popups/popups";

    let mounted: boolean = false;

    onMount(() => {
        mounted = true;

        EventsOn("saved_as", (result: string) => {
            result.includes("error")
                ? showError(result)
                : showSuccess(`${$_("saved_as")} ${result}`);
        });

        EventsOn("enter_password", async () => {
            const { value: password } = await showQuestion(
                $_("enter_password"),
            );
            if (password) {
                EventsEmit("password", password);
            }
        });

        EventsOn("imported_data", (res: string) => {
            // console.log(res);
            if (res === "success") {
                GetLanguage().then((result) => {
                    locale.set(result);
                    EventsEmit(
                        "change_titles",
                        $_("app_title"),
                        $_("select_directory"),
                        $_("select_file"),
                    );
                    showSuccess($_("import_successful"));
                });
            } else {
                // `${res.replace(/^.{1}/g, res[0].toUpperCase())} !!`
                res = res.includes("cannot")
                    ? `${$_("backup_error")}`
                    : `${$_("invalid_password")}`;

                showError(res);
            }
        });
    });

    const showAlert = () => {
        let data: string[] = [
            `${$_("alert_delete_all")}`,
            `${$_("alert_confirm_deleting")}`,
        ];
        showWarning(data).then((result) => {
            if (result.value) {
                Drop().then(() => push("/"));
                showSuccess($_("alert_delete_confirm_msg"));
            }
        });
    };
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex h-screen relative"
    >
        <Language languageName={`🌐 ${$_("language")}`} />

        <section class="m-auto flex flex-col gap-6">
            <h1 class="text-xl font-light m-auto">{$_("settings")}</h1>

            <ul class="flex flex-col items-start">
                <li
                    class="text-warning flex gap-2 justify-between items-center w-full pl-2 rounded-md hover:bg-slate-700"
                >
                    {$_("delete_all_data")}
                    <button
                        class="btn btn-ghost btn-sm"
                        aria-label="Delete All"
                        onclick={showAlert}
                    >
                        <svg
                            fill="#7480ff"
                            class="w-4"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 16 16"
                        >
                            <path
                                d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5M8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5m3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0"
                            />
                        </svg>
                    </button>
                </li>
                <li
                    class="text-success flex gap-2 justify-between items-center w-full pl-2 rounded-md hover:bg-slate-700"
                >
                    {$_("export_data")}
                    <button
                        class="btn btn-ghost btn-sm"
                        aria-label="Export Data"
                        onclick={() => EventsEmit("export_data")}
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="#7480ff"
                            class="w-4"
                            viewBox="0 0 16 16"
                        >
                            <path
                                fill-rule="evenodd"
                                d="M3.5 10a.5.5 0 0 1-.5-.5v-8a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 0 0 1h2A1.5 1.5 0 0 0 14 9.5v-8A1.5 1.5 0 0 0 12.5 0h-9A1.5 1.5 0 0 0 2 1.5v8A1.5 1.5 0 0 0 3.5 11h2a.5.5 0 0 0 0-1z"
                            />
                            <path
                                fill-rule="evenodd"
                                d="M7.646 15.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 14.293V5.5a.5.5 0 0 0-1 0v8.793l-2.146-2.147a.5.5 0 0 0-.708.708z"
                            />
                        </svg>
                    </button>
                </li>
                <li
                    class="text-success flex gap-2 justify-between items-center w-full pl-2 rounded-md hover:bg-slate-700"
                >
                    {$_("import_data")}
                    <button
                        class="btn btn-ghost btn-sm"
                        aria-label="Import Data"
                        onclick={() => EventsEmit("import_data")}
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="#7480ff"
                            class="w-4"
                            viewBox="0 0 16 16"
                        >
                            <path
                                fill-rule="evenodd"
                                d="M3.5 6a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-2a.5.5 0 0 1 0-1h2A1.5 1.5 0 0 1 14 6.5v8a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 14.5v-8A1.5 1.5 0 0 1 3.5 5h2a.5.5 0 0 1 0 1z"
                            />
                            <path
                                fill-rule="evenodd"
                                d="M7.646 11.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 10.293V1.5a.5.5 0 0 0-1 0v8.793L5.354 8.146a.5.5 0 1 0-.708.708z"
                            />
                        </svg>
                    </button>
                </li>
                <li
                    class="text-secondary flex gap-2 justify-between items-center w-full pl-2 rounded-md hover:bg-slate-700"
                >
                    {$_("quit")}
                    <button
                        onclick={() => EventsEmit("quit")}
                        class="btn btn-ghost btn-sm tooltip tooltip-right"
                        data-tip={`${$_("quit")} 'Escape'`}
                        aria-label="Close App"
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="#7480ff"
                            class="w-4"
                            viewBox="0 0 16 16"
                        >
                            <path
                                d="M8.538 1.02a.5.5 0 1 0-.076.998 6 6 0 1 1-6.445 6.444.5.5 0 0 0-.997.076A7 7 0 1 0 8.538 1.02"
                            />
                            <path
                                d="M7.096 7.828a.5.5 0 0 0 .707-.707L2.707 2.025h2.768a.5.5 0 1 0 0-1H1.5a.5.5 0 0 0-.5.5V5.5a.5.5 0 0 0 1 0V2.732z"
                            />
                        </svg>
                    </button>
                </li>
            </ul>
        </section>

        <BackBtn />
    </div>
{/if}
