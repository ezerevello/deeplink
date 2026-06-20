const btnSetup = document.getElementById('btn-setup');
const btnTeardown = document.getElementById('btn-teardown');
const statusText = document.getElementById('status-text');
const lastResponse = document.getElementById('last-response');

async function callApi(path) {
    statusText.textContent = 'working...';
    btnSetup.disabled = true;
    btnTeardown.disabled = true;

    try {
    const res = await fetch(path, { method: 'POST' });
    const txt = await res.text();
    lastResponse.textContent = txt;
    if (!res.ok) {
        statusText.textContent = 'error';
        statusText.className = 'error';
    } else {
        statusText.textContent = 'ok';
        statusText.className = 'success';
    }
    } catch (err) {
    lastResponse.textContent = String(err);
    statusText.textContent = 'network error';
    statusText.className = 'error';
    } finally {
    btnSetup.disabled = false;
    btnTeardown.disabled = false;
    }
}

btnSetup.addEventListener('click', () => callApi('/api/setup'));
btnTeardown.addEventListener('click', () => callApi('/api/teardown'));