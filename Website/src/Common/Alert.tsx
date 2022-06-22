import { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import PropTypes from 'prop-types';

import { alertService, AlertType } from 'Services';


Alert.propTypes = {
    id: PropTypes.string,
    fade: PropTypes.bool
};

Alert.defaultProps = {
    id: 'default-alert',
    fade: true
};

export default function Alert({ id, fade }:{id:string,fade:boolean}) {
    const router = useRouter();
    const [alerts, setAlerts] = useState<any>([]);
    useEffect(() => {
        // subscribe to new alert notifications
        const subscription = alertService.onAlert(id)
            .subscribe(alert => {
                // clear alerts when an empty alert is received
                if (!alert.message) {
                    setAlerts((alerts: any[]) => {
                        // filter out alerts without 'keepAfterRouteChange' flag
                        const filteredAlerts = alerts.filter(x => x.keepAfterRouteChange);
                        // remove 'keepAfterRouteChange' flag on the rest
                        filteredAlerts.forEach(x => delete x.keepAfterRouteChange);
                        return filteredAlerts;
                    });
                } else {
                    // add alert to array
                    setAlerts((alerts: any) => ([...alerts, alert]));

                    // auto close alert if required
                    if (alert.autoClose) {
                        setTimeout(() => removeAlert(alert), 3000);
                    }
                }
            });
        

        // clear alerts on location change
        const onRouteChange = () => alertService.clear(id);
        router.events.on('routeChangeStart', onRouteChange);

        // clean up function that runs when the component unmounts
        return () => {
            // unsubscribe to avoid memory leaks
            subscription.unsubscribe();
            router.events.off('routeChangeStart', onRouteChange);
        };
    }, []);

    function removeAlert(alert: (message?: any) => void) {
        if (fade) {
            // fade out alert
            const alertWithFade = { ...alert, fade: true };
            setAlerts((alerts: any[]) => alerts.map((x: any) => x === alert ? alertWithFade : x));

            // remove alert after faded out
            setTimeout(() => {
                setAlerts((alerts: any[]) => alerts.filter((x: any) => x !== alertWithFade));
            }, 250);
        } else {
            // remove alert
            setAlerts((alerts: any[]) => alerts.filter((x: any) => x !== alert));
        }
    }

    function cssClasses(alert: { type: string | number; fade: any; }) {
        if (!alert) return;

        const classes = ['alert', 'alert-dismissable'];
                
        const alertTypeClass = {
            [AlertType.Success]: 'alert-success',
            [AlertType.Error]: 'alert-danger',
            [AlertType.Info]: 'alert-info',
            [AlertType.Warning]: 'alert-warning'
        }

        classes.push(alertTypeClass[alert.type]);

        if (alert.fade) {
            classes.push('fade');
        }

        return classes.join(' ');
    }

    if (!alerts.length) return null;
    return(
        <div className="container">
            <div className="m-3">
                {alerts.map((alert: any, index: any) =>
                    <div key={index} className={cssClasses(alert) + ' d-flex align-items-center'}>
                        <button type="button" className="btn-close" aria-label="Close" onClick={() => removeAlert(alert)}></button>
                        <span dangerouslySetInnerHTML={{__html: alert.message}}></span>
                    </div>
                )}
            </div>
        </div>
    );
}
