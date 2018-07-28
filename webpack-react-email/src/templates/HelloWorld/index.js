import React from 'react';
import { Box, Item } from 'react-html-email';

import SectionOutline from '../../components/SectionOutline';

export default() => (
    <Box width="100%">
        <SectionOutline>
                <h1>Hello World</h1>
                <p>Hello, this is a wonderful email template</p>
        </SectionOutline>
    </Box>
);