<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import BackBtn from "../lib/BackBtn.svelte";
    import Language from "../lib/Language.svelte";
    import { EventsEmit, EventsOn } from "../../wailsjs/runtime/runtime";
    import { Drop, GetLanguage } from "../../wailsjs/go/main/App";
    import { _, locale } from "svelte-i18n";
    import Swal from "sweetalert2";
    import { push } from "svelte-spa-router";

    let mounted: boolean = false;

    onMount(() => {
        mounted = true;

        EventsOn("saved_as", (file: string) => {
            Swal.fire({
                title: `${$_("saved_as")} ${file}`,
                width: 275,
                icon: "success",
                iconHtml: `
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" viewBox="0 0 16 16">
                    <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/>
                </svg>
                `,
                background: "#1D232A",
                color: "#A6ADBA",
                confirmButtonColor: "#3085d6",
                customClass: {
                    title: "alertTitle",
                    confirmButton: "alertConfirm",
                    icon: "alertIcon",
                },
            });
        });

        EventsOn("enter_password", async () => {
            const { value: password } = await Swal.fire({
                title: `${$_("enter_password")}`,
                width: 275,
                input: "password",
                icon: "question",
                iconHtml: `
                <svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.496 6.033h.825c.138 0 .248-.113.266-.25.09-.656.54-1.134 1.342-1.134.686 0 1.314.343 1.314 1.168 0 .635-.374.927-.965 1.371-.673.489-1.206 1.06-1.168 1.987l.003.217a.25.25 0 0 0 .25.246h.811a.25.25 0 0 0 .25-.25v-.105c0-.718.273-.927 1.01-1.486.609-.463 1.244-.977 1.244-2.056 0-1.511-1.276-2.241-2.673-2.241-1.267 0-2.655.59-2.75 2.286a.237.237 0 0 0 .241.247m2.325 6.443c.61 0 1.029-.394 1.029-.927 0-.552-.42-.94-1.029-.94-.584 0-1.009.388-1.009.94 0 .533.425.927 1.01.927z"/>
                </svg>
                `,
                background: "#1D232A",
                color: "#A6ADBA",
                confirmButtonColor: "#3085d6",
                customClass: {
                    title: "alertTitle",
                    confirmButton: "alertConfirm",
                    icon: "alertIcon",
                    input: "inputDialog",
                },
            });
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

                    Swal.fire({
                        title: `${$_("import_successful")}`,
                        width: 275,
                        icon: "success",
                        iconHtml: `
                        <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/>
                        </svg>
                        `,
                        background: "#1D232A",
                        color: "#A6ADBA",
                        confirmButtonColor: "#3085d6",
                        customClass: {
                            title: "successTitle",
                            confirmButton: "alertConfirm",
                            icon: "alertIcon",
                        },
                    });
                });
            } else {
                // `${res.replace(/^.{1}/g, res[0].toUpperCase())} !!`
                if (res.includes("cannot")) {
                    res = `${$_("backup_error")}`;
                } else {
                    res = `${$_("invalid_password")}`;
                }

                Swal.fire({
                    title: `${res}`,
                    width: 275,
                    icon: "error",
                    iconHtml: `
                    <svg height="32" style="overflow:visible;enable-background:new 0 0 32 32" viewBox="0 0 32 32" width="32"
                        xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                        <g>
                            <g id="Error_1_">
                                <g id="Error">
                                    <circle cx="16" cy="16" id="BG" r="16" style="fill:#D72828;" />
                                    <path d="M14.5,25h3v-3h-3V25z M14.5,6v13h3V6H14.5z" id="Exclamatory_x5F_Sign" style="fill:#E6E6E6;" />
                                </g>
                            </g>
                        </g>
                    </svg>
                    `,
                    background: "#1D232A",
                    color: "#A6ADBA",
                    confirmButtonColor: "#3085d6",
                    customClass: {
                        title: "errorTitle",
                        confirmButton: "alertConfirm",
                        icon: "alertIcon",
                    },
                });
            }
        });
    });

    const showAlert = () =>
        Swal.fire({
            title: `${$_("alert_delete_all")}`,
            width: 275,
            icon: "warning",
            iconHtml: `
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" viewBox="0 0 16 16">
                <path d="M11.46.146A.5.5 0 0 0 11.107 0H4.893a.5.5 0 0 0-.353.146L.146 4.54A.5.5 0 0 0 0 4.893v6.214a.5.5 0 0 0 .146.353l4.394 4.394a.5.5 0 0 0 .353.146h6.214a.5.5 0 0 0 .353-.146l4.394-4.394a.5.5 0 0 0 .146-.353V4.893a.5.5 0 0 0-.146-.353zM8 4c.535 0 .954.462.9.995l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 4.995A.905.905 0 0 1 8 4m.002 6a1 1 0 1 1 0 2 1 1 0 0 1 0-2"/>
            </svg>
            `,
            background: "#1D232A",
            color: "#A6ADBA",
            showCancelButton: true,
            confirmButtonColor: "#3085d6",
            cancelButtonColor: "#d33",
            confirmButtonText: `${$_("alert_confirm_deleting")}`,
            customClass: {
                title: "alertTitle",
                confirmButton: "alertConfirm",
                cancelButton: "alertCancel",
                icon: "alertIcon",
            },
        }).then((result) => {
            if (result.value) {
                Drop().then(() => push("/"));
                Swal.fire({
                    title: `${$_("alert_delete_confirm_msg")}`,
                    width: 275,
                    icon: "success",
                    iconHtml: `
                    <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" fill="currentColor" viewBox="0 0 16 16">
                        <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z"/>
                    </svg>
                    `,
                    background: "#1D232A",
                    color: "#A6ADBA",
                    confirmButtonColor: "#3085d6",
                    customClass: {
                        title: "alertTitle",
                        confirmButton: "alertConfirm",
                        icon: "alertIcon",
                    },
                });
            }
        });
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
