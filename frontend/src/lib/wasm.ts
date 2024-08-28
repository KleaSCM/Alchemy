

export async function initWasm() {

  const wasmModule = await import('../public/wasm/your_module_name.wasm');
  return wasmModule;
}
