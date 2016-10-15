$( document ).ready(function() {

    // Submits the form when button is clicked
    $('#sbmit').click(function() {
      var api = "https://rxuwnhz9hhaobmxlx-mock.stoplight-proxy.io/api/test";
      var site_data = $('#site').val();
      var comprehensive_data = false;
      if($('#comprehensive').is(":checked")) {
        comprehensive_data = true;
      }

      $.getJSON( api, {
        domain: site_data,
        comprehensive: comprehensive_data
      }).done(function(data) {
        render_results(data);
      }).fail(function( jqxhr, textStatus, error ) {
        var err = textStatus + ", " + error;
        console.log( "Request Failed: " + err );
      });
    });

    // Attach the loader to Ajax starts & stops
    var $loader = $('#loader');
    var $info = $('#info');
    var $results = $('#results');
    $(document)
      .ajaxStart(function () {
        $loader.show();
        $results.hide();
        $info.hide();
      })
      .ajaxStop(function () {
        $loader.hide();
        $results.show();
      });

    // Renders results on page according to the JSON data received
    function render_results(jsonData) {
      var data = jsonData;
      var root_domain = data['root-domain'];
      var valid = data['valid'];
      var weaknesses = data['weaknesses'];

      $('#root-domain').text(root_domain);
      if(!valid) {
        $('#results-invalid').show();
        $('#results-valid').hide();
        return;
      }

      // clear table
      $('#results-valid').html('');

      $.each(weaknesses, function(index, weakness) {
        var table = '<table><col width="50%"><col width="50%"><tbody>';
        var main_desc = style(weakness['main-description']);
        var ok_desc = weakness['ok-description'];
        var ok_urls = weakness['ok-urls'];
        var not_ok_desc = weakness['not-ok-description'];
        var not_ok_urls = weakness['not-ok-urls'];

        table += '<tr><td colspan="2" class="scan-desc">' + main_desc + '</td></tr>';
        table += '<tr><th style="text-align:left">Scanned Domains/URLs:</th><th>Scan Results:</th><tr>';

        table += '<tr>';
        table += '<td class="result"><b>' + not_ok_urls.join('<br>') + '</b></td>';
        table += '<td><div class="oracle reason reason-danger">' + not_ok_desc + '</div></td>'
        table += '</tr>';

        table += '<tr>';
        table += '<td class="result"><b>' + ok_urls.join('<br>') + '</b></td>';
        table += '<td><div class="oracle reason reason-safe">' + ok_desc + '</div></td>'
        table += '</tr>';

        table += '</tbody></table>';
        $('#results-valid').append(table);
      });

    }

    function wrap_paragraph(text) {
      return '<p>' + text + '</p>';
    }

    function style(text) {
      var bold_re = /\*([A-z -_\(\)]+)\*/g;
      var bolded_text = text.replace(bold_re, "<b>$1<\/b>");
      var split_bolded_text = bolded_text.split('\n')
      var mapped_bolded_text = split_bolded_text.map(wrap_paragraph);
      return mapped_bolded_text.join('');
    }
});
