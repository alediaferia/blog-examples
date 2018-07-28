
import React, { PureComponent } from 'react';
import PropTypes from 'prop-types';
import { Item } from 'react-html-email';

import './index.scss';

class SectionOutline extends PureComponent {
  static propTypes = {
    children: PropTypes.node.isRequired,
  }
  render() {
    const { children } = this.props;
    const className = 'c-SectionOutline';
    return (
      <Item className={`${className}`} align='center'>
        {children}
      </Item>
    );
  }
}

export default SectionOutline;