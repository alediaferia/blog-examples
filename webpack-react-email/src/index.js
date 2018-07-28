import React, { PureComponent } from 'react';
import ReactDOM from 'react-dom';
import { Box, Item } from 'react-html-email';

/*
 * The EmailContainer will be the body of your Email HTML body.
 * It will receive the right template to load and inject from Webpack
 * and will attempt to load it here and include it in the DOM.
 * 
 * This class expects the EMAIL_TEMPLATE const to be defined.
 */
class EmailContainer extends PureComponent {
  render() {
    // EMAIL_TEMPLATE is defined by the webpack configuration and enables us
    // to include the right template at compile time
    const Template = require(`./templates/${EMAIL_TEMPLATE}`).default;
    return (
      <Box width="600px" height="100%" bgcolor='#f3f3f3' align='center'>
        <Item align='center' valign='top'>
          <Template />
        </Item>
      </Box>
    );
  }
}

ReactDOM.render(<EmailContainer />, document.body);