import Remote from './remote.js';
import DOM from './dom.js';

const initLocal = () => {
    const currentTZ = Intl.DateTimeFormat().resolvedOptions().timeZone
    DOM.setLocalTimezone(currentTZ);

    const dateFormat = Intl.DateTimeFormat('en-US', { 
        dateStyle: 'full', 
        timeStyle: 'medium',
        hour12: false,
        timeZone: currentTZ 
    });
    setInterval(() => { 
        DOM.setLocalTime(dateFormat.format(new Date()));
    }, 100);
};

const initRemote = (remoteTZ) => {
    DOM.setRemoteTimezone(remoteTZ);

    const dateFormat = Intl.DateTimeFormat('en-US', { 
        dateStyle: 'full', 
        timeStyle: 'medium',
        hour12: false,
        timeZone: remoteTZ 
    });
    setInterval(() => { 
        DOM.setRemoteTime(dateFormat.format(new Date()));
    }, 100);
};

Remote.fetchTimeZone().then(tz => initRemote(tz));
window.addEventListener("load", initLocal);
